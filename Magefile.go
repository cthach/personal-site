//go:build mage

package main

import (
	"github.com/magefile/mage/sh"
)

// Compiles the binary
func Build() error {
	return sh.RunV("go", "build", "-o", "build/site", "./cmd")
}

// Executes tests
func Test() error {
	return sh.RunV("go", "test", "./...")
}
