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
// be called from Wren, and to execute Wren code directly from Go.
//
// Foreign Function Limits
//
// Due to Go's inability to generate C-exported functions at runtime, the number of
// foreign methods able to be registered with the Wren VM through this package is limited
// to 128. This number is completely arbitrary, though, and can be changed by modifying
// the directive at the bottom of wren.go and running "go generate". If you feel like
// this number is a terrible default, pull requests will be happily accepted.
//
package wren

// #cgo CFLAGS: -I${SRCDIR}/wren/src/include
// #cgo LDFLAGS: -L${SRCDIR}/wren/lib -lwren -lm
// #include <wren.h>
//
// extern void write(WrenVM*, char*);
// extern void* bindMethod(WrenVM*, char*, char*, bool, char*);
// extern WrenForeignClassMethods bindClass(WrenVM*, char*, char*);
// extern void writeErr(WrenVM*, WrenErrorType, char* module, int line, char* message);
// extern char* loadModule(WrenVM*, char*);
import "C"
import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
	"unsafe"
)

var (
	vmMap     = make(map[*C.WrenVM]*VM)
	errWriter io.Writer
)

// VM is a single instance of a Wren virtual machine.
type VM struct {
	vm               *C.WrenVM
	classes, methods map[string]unsafe.Pointer
	userData         map[string]interface{}
	userDataPtr      unsafe.Pointer
	outWriter        io.Writer
}

// NewVM creates a new Wren virtual machine.
func NewVM() *VM {
	var config C.WrenConfiguration
	C.wrenInitConfiguration(&config)

	config.writeFn = C.WrenWriteFn(C.write)
	config.bindForeignMethodFn = C.WrenBindForeignMethodFn(C.bindMethod)
	config.bindForeignClassFn = C.WrenBindForeignClassFn(C.bindClass)
	config.errorFn = C.WrenErrorFn(C.writeErr)
	config.loadModuleFn = C.WrenLoadModuleFn(C.loadModule)

	vm := VM{vm: C.wrenNewVM(&config)}
	vm.classes = make(map[string]unsafe.Pointer)
	vm.methods = make(map[string]unsafe.Pointer)
	vm.userData = make(map[string]interface{})
	vmMap[vm.vm] = &vm
	runtime.SetFinalizer(&vm, func(vm *VM) {
		C.wrenFreeVM(vm.vm)
		delete(vmMap, vm.vm)
	})

	return &vm
}

// SetModulesDir sets lookup directory for modules to import from.
func (vm *VM) SetModulesDir(path string) {
	vm.setUserData("MODULES_DIR", path)
}

// setUserData preserves (key, val) userdata and makes it available to virtual machine.
func (vm *VM) setUserData(key string, val interface{}) {
	vm.userData[key] = val
	if jval, e := json.Marshal(vm.userData); e == nil {
		if vm.userDataPtr != nil {
			C.free(vm.userDataPtr)
		}
		vm.userDataPtr = unsafe.Pointer(C.CString(string(jval)))
		C.wrenSetUserData(vm.vm, vm.userDataPtr)
	}
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
func (vm *VM) RegisterForeignMethod(fullName string, f interface{}) error {
	ptr, err := registerFunc(fullName, func() {
		if err := handleFunction(vm.vm, f); err != nil {
			panic(err)
		}
	})
	if err != nil {
		return err
	}
	vmMap[vm.vm].methods[fullName] = ptr
	return nil
}

// RegisterForeignClass registers a foreign class with the virtual machine.
func (vm *VM) RegisterForeignClass(className string, f func() interface{}) error {
	ptr, err := registerFunc(className, func() {
		newForeign(vm.vm, f())
	})
	if err != nil {
		return err
	}
	vmMap[vm.vm].classes[className] = ptr
	return nil
}

// SetOutputWriter sets the writer to be used for script output. If this method is never
// called (or called with nil), it uses standard output.
func (vm *VM) SetOutputWriter(w io.Writer) {
	vmMap[vm.vm].outWriter = w
}

// SetErrorWriter sets the writer to be used for script error output. If this method is never
// called (or called with nil), it uses standard error.
func SetErrorWriter(w io.Writer) {
	errWriter = w
}

// GC initiates a garbage collection.
func (vm *VM) GC() {
	C.wrenCollectGarbage(vm.vm)
}

// Interpret interprets the provided Wren source code.
func (vm *VM) Interpret(source string) error {
	c_module := C.CString("main")
	defer C.free(unsafe.Pointer(c_module))
	c_source := C.CString(source)
	defer C.free(unsafe.Pointer(c_source))
	return interpretResultToErr(C.wrenInterpret(vm.vm, c_module, c_source))
}

// InterpretFile interprets the Wren source code in the provided file.
func (vm *VM) InterpretFile(filename string) error {
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	return vm.Interpret(string(contents))
}

// InterpretReader interprets the Wren source code from the provided reader.
// Note that the reader must be read fully before interpretation will begin;
// it's not possible to interpret an infinite stream of input.
func (vm *VM) InterpretReader(r io.Reader) error {
	contents, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	return vm.Interpret(string(contents))
}

// TODO: implement this better. It should automatically pick an available
// slot, then convert the value to something useful.
func (vm *VM) getVariable(module, name string, slot int) {
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

// Value represents a Wren value that Go has a handle to.
type Value struct {
	vm      *C.WrenVM
	value   *C.WrenHandle
	methods map[string]*C.WrenHandle
}

// Variable looks up a variable by name and returns its value.
func (vm *VM) Variable(name string) *Value {
	var (
		c_module = C.CString("main")
		c_name   = C.CString(name)
	)
	defer func() {
		C.free(unsafe.Pointer(c_module))
		C.free(unsafe.Pointer(c_name))
	}()

	C.wrenEnsureSlots(vm.vm, 1)
	C.wrenGetVariable(vm.vm, c_module, c_name, 0)
	value := Value{vm: vm.vm, value: C.wrenGetSlotHandle(vm.vm, 0)}
	if value.value == nil {
		return nil
	}
	value.methods = make(map[string]*C.WrenHandle)
	runtime.SetFinalizer(&value, func(value *Value) {
		for _, method := range value.methods {
			C.wrenReleaseHandle(vm.vm, method)
		}
		C.wrenReleaseHandle(vm.vm, value.value)
	})
	return &value
}

// Call calls the method with the given signature that belongs to the given value.
//
// The receiver should be the value on which the method is defined; a class reference
// for static methods, and an instance of a class for instance methods. The signature
// is a standard Wren method signature, and any parameters it expects will follow.
func (v *Value) Call(signature string, params ...interface{}) (interface{}, error) {
	f := v.methods[signature]
	if f == nil {
		c_signature := C.CString(signature)
		defer C.free(unsafe.Pointer(c_signature))
		f = C.wrenMakeCallHandle(v.vm, c_signature)
		v.methods[signature] = f
	}
	C.wrenEnsureSlots(v.vm, C.int(len(params)+1))
	C.wrenSetSlotHandle(v.vm, 0, v.value)
	for i, param := range params {
		saveToSlot(v.vm, i+1, reflect.ValueOf(param))
	}
	if err := interpretResultToErr(C.wrenCall(v.vm, f)); err != nil {
		return nil, err
	}
	if retval := getFromSlot(v.vm, 0, nil); retval.IsValid() {
		return retval.Interface(), nil
	}
	return nil, nil
}

// newForeign allocates a new foreign object.
//
// This method should only be called from a foreign class allocation function.
// It takes an instance of the VM and a newly allocated foreign object ("foreign"
// meaning that it's created in Go and not Wren) and makes it available to Wren.
func newForeign(vm *C.WrenVM, x interface{}) {
	var (
		v   = reflect.Indirect(reflect.ValueOf(x))
		t   = v.Type()
		ptr = C.wrenSetSlotNewForeign(vm, C.int(0), C.int(0), C.size_t(t.Size()))
	)
	reflect.NewAt(t, ptr).Elem().Set(v)
}

// handleFunction is a helper method for foreign methods.
//
// This method takes two parameters: a reference to the virtual machine instance
// (which should be the only parameter provided in the C-exported callback)
// and a Go function. The function's signature must match the one expected by Wren.
// If it doesn't, this call will return an error, but the call to Interpret() will not.
//
// For examples, check out the test package.
func handleFunction(vm *C.WrenVM, f interface{}) (err error) {
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
		slot := i + offset

		// If the receiver value is inaccessible from C, it likely just means that
		// it's a native class with a foreign method. Rather than panic, we simply
		// advance to the first parameter and continue from there.
		if i == 0 && C.wrenGetSlotType(vm, C.int(slot)) == C.WREN_TYPE_UNKNOWN {
			offset++
			slot++
		}

		it := ft.In(i)
		params[i] = getFromSlot(vm, slot, &it)
	}

	returnValues := fv.Call(params)
	// TODO: allow returning a second value if it's an `error`, like the template packages
	if len(returnValues) == 1 {
		saveToSlot(vm, 0, returnValues[0])
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

//export loadModule
func loadModule(vm *C.WrenVM, name *C.char) *C.char {
	var module string = C.GoString(name)
	var source string

	// Ensure module does not have undesired characters
	// that can pose thread to remote-code-inclusions
	if !strings.Contains(module, "..") {
		// Proceed to load from the configured modules directory only
		var jvalPtr unsafe.Pointer = C.wrenGetUserData(vm)
		if jvalPtr != nil {
			userData := make(map[string]interface{})
			jval := C.GoString((*C.char)(jvalPtr))
			if e := json.Unmarshal([]byte(jval), &userData); e == nil {
				if modulesDir, ok := userData["MODULES_DIR"]; ok {
					// Precedence (modules_dir/module_name.wren) next (modules_dir/module_name/module.wren)
					if fdata, e := ioutil.ReadFile(filepath.Join(modulesDir.(string), module) + ".wren"); e == nil {
						source = string(fdata)
					} else if fdata, e = ioutil.ReadFile(filepath.Join(modulesDir.(string), module, "module.wren")); e == nil {
						source = string(fdata)
					}
				}
			}
		}
	}

	return C.CString(source)
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
func writeErr(vm *C.WrenVM, errorType C.WrenErrorType, module *C.char, line C.int, message *C.char) {
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

func interpretResultToErr(result C.WrenInterpretResult) error {
	switch result {
	case C.WREN_RESULT_SUCCESS:
		return nil

	case C.WREN_RESULT_COMPILE_ERROR:
		return errors.New("compilation error")

	case C.WREN_RESULT_RUNTIME_ERROR:
		return errors.New("runtime error")

	default:
		panic("unreachable")
	}
}

func saveToSlot(vm *C.WrenVM, slot int, v reflect.Value) {
	c_slot := C.int(slot)
	switch v.Kind() {
	case reflect.Bool:
		c_value := C.bool(v.Interface().(bool))
		C.wrenSetSlotBool(vm, c_slot, c_value)

	case reflect.Float32, reflect.Float64:
		c_value := C.double(v.Float())
		C.wrenSetSlotDouble(vm, c_slot, c_value)

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		c_value := C.double(v.Int())
		C.wrenSetSlotDouble(vm, c_slot, c_value)

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		c_value := C.double(v.Uint())
		C.wrenSetSlotDouble(vm, c_slot, c_value)

	case reflect.String:
		c_value := C.CString(v.Interface().(string))
		defer C.free(unsafe.Pointer(c_value))
		C.wrenSetSlotString(vm, c_slot, c_value)

	default:
		panic(fmt.Sprintf("don't know how to save this to a slot: %s", v.Type().Name()))
	}
}

func getFromSlot(vm *C.WrenVM, slot int, in *reflect.Type) reflect.Value {
	c_slot := C.int(slot)
	switch C.wrenGetSlotType(vm, c_slot) {
	case C.WREN_TYPE_BOOL:
		return reflect.ValueOf(bool(C.wrenGetSlotBool(vm, c_slot)))

	case C.WREN_TYPE_NUM:
		n := reflect.ValueOf(float64(C.wrenGetSlotDouble(vm, c_slot)))
		if in != nil {
			return n.Convert(*in)
		}
		return n

	case C.WREN_TYPE_FOREIGN:
		if in == nil {
			panic("can't return foreign value without type information!")
		}
		ptr := C.wrenGetSlotForeign(vm, c_slot)
		return reflect.NewAt((*in).Elem(), ptr)

	case C.WREN_TYPE_LIST:
		panic("not sure how to get a list value from the slot")

	case C.WREN_TYPE_NULL:
		return reflect.Value{}

	case C.WREN_TYPE_STRING:
		str := C.GoString(C.wrenGetSlotString(vm, c_slot))
		return reflect.ValueOf(str)

	case C.WREN_TYPE_UNKNOWN:
		panic(fmt.Sprintf("received an inaccessible-from-C parameter in slot %d", slot))

	default:
		panic("unreachable")
	}
}

// Change 128 to a different number to enable more foreign class/method registrations.
//go:generate go run cgluer.go 128
