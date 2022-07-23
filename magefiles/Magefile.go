package main

import (
	"os"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

// Removes all artifacts
func Clean() error {
	return os.RemoveAll("build")
}

type Build mg.Namespace

func (Build) Container() error {
	return sh.RunV(
		"podman",
		"build",
		"-t", "docker.io/cthach/personal-site:latest",
		".",
	)
}

// Compiles the binary
func (Build) Binary() error {
	return sh.RunV("go", "build", "-o", "build/site", "./cmd")
}

// Executes tests
func Test() error {
	return sh.RunV("go", "test", "./...")
}
