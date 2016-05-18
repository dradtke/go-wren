// +build ignore

// This file is intended to be executed by "go generate".

package main

import (
	"os"
	"strconv"
	"text/template"
)

var fileTemplate = template.Must(template.New("").Parse(`package wren

/*
{{range .}}extern void f{{.}}(void* vm);
{{end}}
static inline void* get_f(int i) {
	switch (i) {
		{{range .}}case {{.}}: return f{{.}};
		{{end}}default: return (void*)(0);
	}
}
*/
import "C"
import (
	"errors"
	"sync"
	"unsafe"
)

const MAX_REGISTRATIONS = {{len .}}

var (
	fMap = make(map[int]func())
	fMapGuard sync.Mutex
	counter int
)

{{range .}}
//export f{{.}}
func f{{.}}(vm unsafe.Pointer) {
	f := fMap[{{.}}]
	if f == nil {
		panic("function {{.}} not registered")
	}
	f()
}
{{end}}

func registerFunc(name string, f func()) (unsafe.Pointer, error) {
	if (counter+1) >= MAX_REGISTRATIONS {
		return nil, errors.New("maximum function registration reached")
	}

	fMapGuard.Lock()
	defer fMapGuard.Unlock()

	fMap[counter] = f
	ptr := C.get_f(C.int(counter))
	counter++
	return ptr, nil
}
`))

func main() {
	if len(os.Args) == 1 {
		panic("no number provided")
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic(err)
	}

	f, err := os.Create("cglue.go")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	data := make([]int, n)
	for i := 0; i < n; i++ {
		data[i] = i
	}
	if err := fileTemplate.Execute(f, data); err != nil {
		panic(err)
	}
}
