// +build mage

package main

import (
	"github.com/magefile/mage/sh"
)

// Cleans up build directories
func Clean() {
	sh.Rm("build")
}
