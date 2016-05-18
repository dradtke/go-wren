package wren_test

import (
	"fmt"

	"github.com/dradtke/go-wren"
)

// Foreign classes can be any type, but usually you'll want a struct.
type God struct {
	msg string
}

// The constructor for a foreign type takes no arguments and returns
// an interface{} value representing the new object.
func NewGod() interface{} {
	return &God{msg: "Do my bidding, %s!"}
}

// Foreign methods take the receiver as its first parameter.
func GetGodsMessage(g *God, name string) string {
	return fmt.Sprintf(g.msg, name)
}

func Example_foreignClass() {
	// A simple program that constructs an instance of a foreign class
	// and calls a foreign method on it.
	const program = `
		foreign class God {
			construct new() {}
			foreign getMessage(name)
		}

		var god = God.new()
		System.print(god.getMessage("Damien"))
	`

	// Initialize the virtual machine and register the foreign class/method.
	vm := wren.NewVM()
	vm.RegisterForeignClass("God", NewGod)
	vm.RegisterForeignMethod("God.getMessage(_)", GetGodsMessage)

	if err := vm.Interpret(program); err != nil {
		panic(err)
	}
}
