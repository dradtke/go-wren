package wren_test

import (
	"fmt"

	"github.com/dradtke/go-wren"
)

func Example_callWren() {
	const program = `
		class Bird {
			static fly(where) {
				return "Flying to %(where)!"
			}
		}
	`

	vm := wren.NewVM()
	if err := vm.Interpret(program); err != nil {
		panic(err)
	}

	response, err := vm.Variable("Bird").Call("fly(_)", "Chicago")
	if err != nil {
		panic(err)
	}

	fmt.Println(response)
}
