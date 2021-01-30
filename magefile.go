// +build mage

package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

const name string = "boltchat"
const buildDir string = "build"

const serverPrefix string = "server"
const clientPrefix string = "client"

const serverEntry string = "cmd/server/server.go"
const clientEntry string = "cmd/client/client.go"

type (
	Build  mg.Namespace
	Test   mg.Namespace
	Docker mg.Namespace
	CI     mg.Namespace
)

type BuildOptions struct {
	Static    bool
	Extension string
	Prefix    string
}

func build(os string, arch string, entry string, opts BuildOptions) error {
	env := map[string]string{
		"GOOS":   os,
		"GOARCH": arch,
	}

	// Build static binary
	if opts.Static {
		env["CGO_ENABLED"] = "0"
	}

	outputName := fmt.Sprintf(
		"%s-%s-%s-%s", name, opts.Prefix, os, arch,
	)

	outputPath := path.Join(
		buildDir,
		outputName,
	)

	if opts.Extension != "" {
		outputPath += fmt.Sprintf(".%s", opts.Extension)
	}

	args := []string{
		"build",
		"-o",
		outputPath,
		"-ldflags",
		"-s -w",
		entry,
	}

	fmt.Println(args)

	return sh.RunWith(
		env, "go", args...,
	)
}

/*
Build
*/

// Builds all binaries
func (Build) All() {
	mg.Deps(
		Build.ServerDarwinAmd64,
		Build.ServerLinuxAmd64,
		Build.ServerWindowsAmd64,

		Build.ClientDarwinAmd64,
		Build.ClientLinuxAmd64,
		Build.ClientWindowsAmd64,
	)
}

// Builds the server binary for Linux (amd64)
func (Build) ServerLinuxAmd64() error {
	return build("linux", "amd64", serverEntry, BuildOptions{Prefix: serverPrefix})
}

// Builds the server binary for Windows (amd64)
func (Build) ServerWindowsAmd64() error {
	return build("windows", "amd64", serverEntry, BuildOptions{
		Extension: "exe",
		Prefix:    serverPrefix,
	})
}

// Builds the server binary for Darwin/macOS (amd64)
func (Build) ServerDarwinAmd64() error {
	return build("darwin", "amd64", serverEntry, BuildOptions{Prefix: serverPrefix})
}

// Builds the server binary for Darwin/macOS (arm64, M1)
// func (Build) ServerDarwinArm64() error {
// 	return build("darwin", "arm64", serverEntry, false)
// }

// Builds the static server binary for use in a Docker container
func (Build) ServerStatic() error {
	return build("linux", "amd64", serverEntry, BuildOptions{
		Static: true,
		Prefix: serverPrefix,
	})
}

// Builds the client binary for Linux (amd64)
func (Build) ClientLinuxAmd64() error {
	return build("linux", "amd64", clientEntry, BuildOptions{Prefix: clientPrefix})
}

// Builds the client binary for Windows (amd64)
func (Build) ClientWindowsAmd64() error {
	return build("windows", "amd64", clientEntry, BuildOptions{
		Extension: "exe",
		Prefix:    clientPrefix,
	})
}

// Builds the client binary for Darwin/macOS (amd64)
func (Build) ClientDarwinAmd64() error {
	return build("darwin", "amd64", clientEntry, BuildOptions{Prefix: clientPrefix})
}

/*
Test
*/

// Runs all unit tests
func (Test) Unit() error {
	return sh.RunV("go", "test", "-v", "./...")
}

/*
Docker
*/

// Builds a Docker image for the server
func (Docker) Build() error {
	return sh.RunV("docker", "build", ".", "-t", name)
}

/*
CI/CD
*/

// Compresses all binaries into a single tarball
func (CI) CompressBinaries() error {
	return sh.Run("tar", "-cvzf", "binaries.tar.gz", "build")
}

/*
Misc
*/

// Adds license headers to source files
func License() {
	paths := make([]string, 0)

	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if strings.HasSuffix(path, ".go") {
			paths = append(paths, path)
		}
		return nil
	})

	sh.Run(
		"addlicense",
		"-l", "apache",
		"-c", "The boltchat Authors",
		"client", "server", "protocol", "cmd", "util",
	)
}

// Cleans up build directories
func Clean() {
	sh.Rm("build")
}
