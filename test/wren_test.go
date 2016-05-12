package test

import (
	"testing"
)

func TestOutput(t *testing.T) {
	testOutput(t)
}

func TestCompilationError(t *testing.T) {
	testCompilationError(t)
}

func TestForeignClass(t *testing.T) {
	testForeignClass(t)
}

func TestForeignMethod(t *testing.T) {
	testForeignMethod(t)
}
