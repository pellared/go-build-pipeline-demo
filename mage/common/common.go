package common

import "github.com/magefile/mage/sh"

// Fmt runs go fmt.
func Fmt() error {
	return sh.Run("go", "fmt", "./...")
}
