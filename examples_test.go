package wren_test

import (
	"github.com/dradtke/go-wren"
)

func ExampleWrenVM_Interpret() {
	const program = `
		System.print("Hello from Wren!")
	`

	vm := wren.NewVM()
	if err := vm.Interpret(program); err != nil {
		panic(err)
	}
}
