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
	return sh.RunWithV(
		map[string]string{
			"CGO_ENABLED": "0",
			"GOOS":        "linux",
			"GOARCH":      "amd64",
		},
		"go",
		"build",
		"-a",
		"-ldflags", `-w -extldflags "-static"`,
		"-o", "build/site",
		"./cmd",
	)
}

// Executes tests
func Test() error {
	return sh.RunV("go", "test", "./...")
}
