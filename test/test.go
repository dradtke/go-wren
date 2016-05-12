package test

// extern void callPerson(void*);
//
// extern void newGod(void*);
// extern void getGodsMessage(void*);
//
// extern void goMathAdd(void*);
import "C"
import (
	"bytes"
	"fmt"
	"io/ioutil"
	"testing"
	"unsafe"

	"github.com/dradtke/go-wren"
)

// Redirect VM output to a buffer, and ensure that it contains the expected value.
func testOutput(t *testing.T) {
	var (
		buf bytes.Buffer
		vm  = wren.NewVM()
	)
	vm.SetOutputWriter(&buf)

	if err := vm.Interpret(`System.print("Hello, Wren!")`); err != nil {
		t.Log("interpretation error: ", err)
		t.FailNow()
	}
	if buf.String() != "Hello, Wren!\n" {
		t.Errorf("unexpected output: %s", buf.String())
	}
}

// Force an error out of the VM.
func testCompilationError(t *testing.T) {
	vm := wren.NewVM()
	wren.SetErrorWriter(ioutil.Discard)

	if err := vm.Interpret(`Don't mind me!`); err == nil {
		t.Error("interpretation of invalid program failed to return an error")
	}
}

// Register a foreign class and call one of its methods.
func testForeignClass(t *testing.T) {
	var buf bytes.Buffer

	vm := wren.NewVM()
	vm.SetOutputWriter(&buf)
	vm.RegisterForeignClass("God", C.newGod)
	vm.RegisterForeignMethod("God.getMessage(_)", C.getGodsMessage)

	if err := vm.Interpret(`
		foreign class God {
			construct new() {}
			foreign getMessage(name)
		}

		var god = God.new()
		System.print(god.getMessage("Damien"))
	`); err != nil {
		t.Log(err.Error())
		t.FailNow()
	}

	if buf.String() != "Do my bidding, Damien!\n" {
		t.Errorf("unexpected output: %s", buf.String())
	}
}

// Register a static foreign method and ensure that it gives the
// expected response.
func testForeignMethod(t *testing.T) {
	var buf bytes.Buffer

	vm := wren.NewVM()
	vm.SetOutputWriter(&buf)
	vm.RegisterForeignMethod("static GoMath.add(_,_)", C.goMathAdd)

	if err := vm.Interpret(`
		class GoMath {
			foreign static add(x, y)
		}

		System.write(GoMath.add(2, 3))
	`); err != nil {
		t.Log(err.Error())
		t.FailNow()
	}

	if buf.String() != "5" {
		t.Errorf("unexpected output: %s", buf.String())
	}

	defer func() {
		if r := recover(); r == nil {
			t.Error("GoMath.add(_,_) call succeeded with invalid parameters")
		}
	}()

	// This call should panic.
	vm.Interpret(`GoMath.add("x", "y")`)
}

type god struct {
	msg string
}

//export newGod
func newGod(vm unsafe.Pointer) {
	wren.NewForeign(vm, &god{msg: "Do my bidding, %s!"})
}

//export getGodsMessage
func getGodsMessage(vm unsafe.Pointer) {
	f := func(g *god, name string) string {
		return fmt.Sprintf(g.msg, name)
	}

	if err := wren.HandleFunction(vm, f); err != nil {
		panic(err)
	}
}

//export goMathAdd
func goMathAdd(vm unsafe.Pointer) {
	f := func(x, y int) int {
		return x + y
	}

	if err := wren.HandleFunction(vm, f); err != nil {
		panic(err)
	}
}
