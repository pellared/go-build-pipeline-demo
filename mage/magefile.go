// +build mage

package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

// All runs build pipeline.
func All() {
	mg.SerialDeps(Clean, Fmt, Test)
}

// Clean remove files created during build.
func Clean() error {
	files, err := filepath.Glob("coverage.*")
	if err != nil {
		return fmt.Errorf("glob failed: %v", err)
	}

	for _, file := range files {
		if cErr := os.Remove(file); cErr != nil {
			err = cErr
			fmt.Println("failed to remove", file, err)
			continue
		}
		if mg.Verbose() {
			fmt.Println("removed", file)
		}
	}
	return err
}

// Fmt runs go fmt.
func Fmt() error {
	return sh.Run("go", "fmt", "./...")
}

// Test runs go test with race detector and code covarage.
func Test() error {
	err1 := sh.Run("go", "test", "-race", "-covermode=atomic", "-coverprofile=coverage.out", "./...")
	err2 := sh.Run("go", "tool", "cover", "-html=coverage.out", "-o", "coverage.html")
	if err1 != nil {
		return err1
	}
	if err2 != nil {
		return err2
	}
	return nil
}
