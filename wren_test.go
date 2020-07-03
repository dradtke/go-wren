package wren_test

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/dradtke/go-wren"
)

func TestCompilationError(t *testing.T) {
	vm := wren.NewVM()
	wren.SetErrorWriter(ioutil.Discard)

	if err := vm.Interpret(`Don't mind me, I'm just an invalid Wren program!`); err == nil {
		t.Error("interpretation of invalid program failed to return an error")
	}
}

func TestOutputRedirect(t *testing.T) {
	var buf bytes.Buffer
	vm := wren.NewVM()
	vm.SetOutputWriter(&buf)

	if err := vm.Interpret(`System.print("Hello, Wren!")`); err != nil {
		t.Log("interpretation error: ", err)
		t.FailNow()
	}
	if buf.String() != "Hello, Wren!\n" {
		t.Errorf("unexpected output: %s", buf.String())
	}
}

func TestForeignMethod(t *testing.T) {
	var buf bytes.Buffer
	vm := wren.NewVM()
	vm.SetOutputWriter(&buf)

	vm.RegisterForeignMethod("static GoMath.add(_,_)", func(a, b int) int {
		return a + b
	})

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

func TestForeignClass(t *testing.T) {
	type God struct {
		msg string
	}

	var buf bytes.Buffer
	vm := wren.NewVM()
	vm.SetOutputWriter(&buf)

	vm.RegisterForeignClass("God", func() interface{} {
		return &God{msg: "Do my bidding, %s!"}
	})

	vm.RegisterForeignMethod("God.getMessage(_)", func(g *God, name string) string {
		return fmt.Sprintf(g.msg, name)
	})

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

func TestCallWren(t *testing.T) {
	vm := wren.NewVM()

	if err := vm.Interpret(`
		class WrenMath {
			static do_add(a, b) {
				return a + b
			}
		}
	`); err != nil {
		t.Log(err.Error())
		t.FailNow()
	}

	x, err := vm.Variable("WrenMath").Call("do_add(_,_)", 2, 3)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	n, ok := x.(float64)
	if !ok {
		t.Logf("WrenMath.add(2, 3) returned unexpected value: %v", x)
		t.FailNow()
	}

	if n != 5 {
		t.Errorf("WrenMath.add(2, 3) returned unexpected value: %v", x)
	}
}

func TestLoadModule(t *testing.T) {
	vm := wren.NewVM()
	wren.SetModulesDir("test_modules")

	if err := vm.Interpret(`import "hello" for Hello
		Hello.World()`); err != nil {
		t.Log("module load error: ", err)
		t.FailNow()
	}
}
