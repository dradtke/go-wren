// Package wren provides bindings to the Wren scripting language.
//
// (To learn more about Wren, visit http://wren.io/index.html)
//
// This package facilitates running Wren code from Go. At its simplest,
// all you need to do is create a new virtual machine instance and interpret
// some Wren code:
//
//      package main
//
//      import (
//      	"github.com/dradtke/go-wren"
//      	"log"
//      )
//
//      func main() {
//      	vm := wren.NewVM()
//      	if err := vm.Interpret(`System.print("Hello, Wren!")`); err != nil {
//      		log.Println(err)
//      	}
//      }
//
// However, it's also possible to register foreign classes and methods in Go that can
// be called from Wren. The reverse is not yet possible, but it likely will be at
// some point.
//
// For more usage examples, check out the test package.
//
package wren

// Tips: https://github.com/golang/go/wiki/cgo

// #cgo CFLAGS: -I${SRCDIR}/wren/src/include
// #cgo LDFLAGS: -L${SRCDIR}/wren/lib -lwren
// #include <wren.h>
//
// extern void write(WrenVM*, char*);
// extern void* bindMethod(WrenVM*, char*, char*, bool, char*);
// extern WrenForeignClassMethods bindClass(WrenVM*, char*, char*);
// extern void writeErr(WrenErrorType, char* module, int line, char* message);
import "C"
import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"reflect"
	"runtime"
	"unsafe"
)

var (
	vmMap     = make(map[*C.WrenVM]*WrenVM)
	errWriter io.Writer
)

// WrenVM is a single instance of a Wren virtual machine.
type WrenVM struct {
	vm               *C.WrenVM
	classes, methods map[string]unsafe.Pointer
	outWriter        io.Writer
}

// NewVM creates a new Wren virtual machine.
func NewVM() *WrenVM {
	var config C.WrenConfiguration
	C.wrenInitConfiguration(&config)

	config.writeFn = C.WrenWriteFn(C.write)
	config.bindForeignMethodFn = C.WrenBindForeignMethodFn(C.bindMethod)
	config.bindForeignClassFn = C.WrenBindForeignClassFn(C.bindClass)
	config.errorFn = C.WrenErrorFn(C.writeErr)

	vm := WrenVM{vm: C.wrenNewVM(&config)}
	vm.classes = make(map[string]unsafe.Pointer)
	vm.methods = make(map[string]unsafe.Pointer)
	vmMap[vm.vm] = &vm
	runtime.SetFinalizer(&vm, func(vm *WrenVM) {
		C.wrenFreeVM(vm.vm)
		delete(vmMap, vm.vm)
	})
	return &vm
}

// RegisterForeignMethod registers a foreign method with the virtual machine.
//
// fullName should be a fully-qualified description string for the method. In particular,
// it should look like this:
//
//     "[static ]<class>.<method>"
//
// At minimum, it should have the class name and the method name separated by a period,
// optionally with the word "static" out front to denote that it's a static method.
func (vm *WrenVM) RegisterForeignMethod(fullName string, x interface{}) {
	vmMap[vm.vm].methods[fullName] = unsafe.Pointer(reflect.ValueOf(x).Pointer())
}

// RegisterForeignClass registers a foreign class with the virtual machine.
func (vm *WrenVM) RegisterForeignClass(className string, x interface{}) {
	vmMap[vm.vm].classes[className] = unsafe.Pointer(reflect.ValueOf(x).Pointer())
}

// SetOutputWriter sets the writer to be used for script output. If this method is never
// called (or called with nil), it uses standard output.
func (vm *WrenVM) SetOutputWriter(w io.Writer) {
	vmMap[vm.vm].outWriter = w
}

// SetErrorWriter sets the writer to be used for script error output. If this method is never
// called (or called with nil), it uses standard error.
func SetErrorWriter(w io.Writer) {
	errWriter = w
}

// GC initiates a garbage collection.
func (vm *WrenVM) GC() {
	C.wrenCollectGarbage(vm.vm)
}

// Interpret interprets the provided Wren source code.
func (vm *WrenVM) Interpret(source string) error {
	c_source := C.CString(source)
	defer C.free(unsafe.Pointer(c_source))
	switch C.wrenInterpret(vm.vm, c_source) {
	case C.WREN_RESULT_SUCCESS:
		return nil

	case C.WREN_RESULT_COMPILE_ERROR:
		return errors.New("compile error")

	case C.WREN_RESULT_RUNTIME_ERROR:
		return errors.New("runtime error")

	default:
		panic("unreachable")
	}
}

// InterpretFile interprets the Wren source code in the provided file.
func (vm *WrenVM) InterpretFile(filename string) error {
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	return vm.Interpret(string(contents))
}

// InterpretReader interprets the Wren source code from the provided reader.
// Note that the reader must be read fully before interpretation will begin;
// it's not possible to interpret an infinite stream of input.
func (vm *WrenVM) InterpretReader(r io.Reader) error {
	contents, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	return vm.Interpret(string(contents))
}

// NewForeign allocates a new foreign object.
//
// This method should only be called from a foreign class allocation function.
// It takes an instance of the VM and a newly allocated foreign object ("foreign"
// meaning that it's created in Go and not Wren) and makes it available to Wren.
func NewForeign(vm unsafe.Pointer, x interface{}) {
	var (
		v   = reflect.Indirect(reflect.ValueOf(x))
		t   = v.Type()
		ptr = C.wrenSetSlotNewForeign(vm, C.int(0), C.int(0), C.size_t(t.Size()))
	)
	reflect.NewAt(t, ptr).Elem().Set(v)
}

// TODO: implement this better. It should automatically pick an available
// slot, then convert the value to something useful.
func (vm *WrenVM) getVariable(module, name string, slot int) {
	var (
		c_module = C.CString(module)
		c_name   = C.CString(name)
	)
	defer func() {
		C.free(unsafe.Pointer(c_module))
		C.free(unsafe.Pointer(c_name))
	}()
	C.wrenGetVariable(vm.vm, c_module, c_name, C.int(slot))
}

// HandleFunction is a helper method for foreign methods.
//
// This method takes two parameters: a reference to the virtual machine instance
// (which should be the only parameter provided in the C-exported callback)
// and a Go function. The function's signature must match the one expected by Wren.
// If it doesn't, this call will return an error, but the call to Interpret() will not.
//
// For examples, check out the test package.
func HandleFunction(vm unsafe.Pointer, f interface{}) (err error) {
	defer func() {
		if r := recover(); r != nil {
			// Fuck.
			switch x := r.(type) {
			case error:
				err = x
			case string:
				err = errors.New(x)
			default:
				err = fmt.Errorf("%v", x)
			}
		}
	}()

	var (
		fv     = reflect.ValueOf(f)
		ft     = fv.Type()
		params = make([]reflect.Value, ft.NumIn())
	)

	var offset int
	for i := 0; i < ft.NumIn(); i++ {
		slot := C.int(i + offset)

		// If the receiver value is inaccessible from C, it likely just means that
		// it's a native class with a foreign method. Rather than panic, we simply
		// advance to the first parameter and continue from there.
		if i == 0 && C.wrenGetSlotType(vm, slot) == C.WREN_TYPE_UNKNOWN {
			offset++
			slot++
		}

		switch C.wrenGetSlotType(vm, slot) {
		case C.WREN_TYPE_BOOL:
			params[i] = reflect.ValueOf(bool(C.wrenGetSlotBool(vm, slot)))

		case C.WREN_TYPE_NUM:
			n := float64(C.wrenGetSlotDouble(vm, slot))
			params[i] = reflect.ValueOf(n).Convert(ft.In(i))

		case C.WREN_TYPE_FOREIGN:
			ptr := C.wrenGetSlotForeign(vm, slot)
			params[i] = reflect.NewAt(ft.In(i).Elem(), ptr)

		case C.WREN_TYPE_LIST:
			panic("not sure how to get a list value from the slot")

		case C.WREN_TYPE_NULL:
			params[i] = reflect.ValueOf(nil)

		case C.WREN_TYPE_STRING:
			str := C.GoString(C.wrenGetSlotString(vm, slot))
			params[i] = reflect.ValueOf(str)

		case C.WREN_TYPE_UNKNOWN:
			panic(fmt.Sprintf("received an inaccessible-from-C parameter in slot %d", slot))
		}
	}

	returnValues := fv.Call(params)
	// TODO: allow returning a second value if it's an `error`
	if len(returnValues) == 1 {
		slot := C.int(0)
		switch returnValues[0].Kind() {
		case reflect.Bool:
			c_value := C.bool(returnValues[0].Interface().(bool))
			C.wrenSetSlotBool(vm, slot, c_value)

		case reflect.Float32, reflect.Float64:
			c_value := C.double(returnValues[0].Float())
			C.wrenSetSlotDouble(vm, slot, c_value)

		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			c_value := C.double(returnValues[0].Int())
			C.wrenSetSlotDouble(vm, slot, c_value)

		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			c_value := C.double(returnValues[0].Uint())
			C.wrenSetSlotDouble(vm, slot, c_value)

		case reflect.String:
			c_value := C.CString(returnValues[0].Interface().(string))
			defer C.free(unsafe.Pointer(c_value))
			C.wrenSetSlotString(vm, slot, c_value)

		default:
			panic(fmt.Sprintf("don't know how to return value of type: %s", returnValues[0].Type().Name()))
		}
	}
	return
}

//export write
func write(vm *C.WrenVM, text *C.char) {
	out := vmMap[vm].outWriter
	if out == nil {
		out = os.Stdout
	}
	fmt.Fprint(out, C.GoString(text))
}

//export bindMethod
func bindMethod(vm *C.WrenVM, c_module, c_className *C.char, c_isStatic C.bool, c_signature *C.char) unsafe.Pointer {
	module := C.GoString(c_module)
	if module != "main" {
		return unsafe.Pointer(nil)
	}

	var (
		className = C.GoString(c_className)
		isStatic  = bool(c_isStatic)
		signature = C.GoString(c_signature)
		fullName  bytes.Buffer
	)

	if isStatic {
		fullName.WriteString("static ")
	}
	fullName.WriteString(className)
	fullName.WriteString(".")
	fullName.WriteString(signature)

	if f, ok := vmMap[vm].methods[fullName.String()]; ok {
		return f
	}
	return unsafe.Pointer(nil)
}

//export bindClass
func bindClass(vm *C.WrenVM, c_module, c_className *C.char) C.WrenForeignClassMethods {
	module := C.GoString(c_module)
	if module != "main" {
		panic("tried to bind foreign class from non-main module")
	}

	className := C.GoString(c_className)
	if c, ok := vmMap[vm].classes[className]; ok {
		// Might be a good idea to support finalizers, but since this is Go,
		// I don't think they're actually necessary.
		return C.WrenForeignClassMethods{
			allocate: C.WrenForeignMethodFn(c),
			finalize: nil,
		}
	}

	panic(fmt.Sprintf("foreign class %s not found", className))
}

//export writeErr
func writeErr(errorType C.WrenErrorType, module *C.char, line C.int, message *C.char) {
	out := errWriter
	if out == nil {
		out = os.Stderr
	}

	switch errorType {
	case C.WREN_ERROR_COMPILE:
		fmt.Fprintf(out, "compilation error: %s:%d: %s\n", C.GoString(module), int(line), C.GoString(message))

	case C.WREN_ERROR_RUNTIME:
		fmt.Fprintf(out, "runtime error: %s", C.GoString(message))

	case C.WREN_ERROR_STACK_TRACE:
		fmt.Fprintf(out, "\t%s:%d: %s\n", C.GoString(module), int(line), C.GoString(message))

	default:
		panic("impossible error type")
	}
}
