// Copyright 2021 The boltchat Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package errs

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

type SyntaxError struct {
	Error error
	Desc  string
}

/*
Emerg displays a message to the user that something has
gone wrong internally, and immediately stops program
execution afterwards.
*/
func Emerg(err error) {
	fmt.Printf(color.HiRedString(
		"An unexpected error has occurred.\nPlease consider creating " +
			"an issue at <https://github.com/bolt-chat/boltchat/issues> " +
			"if this is repetitive behaviour.\n",
	))

	// Immediately stop execution
	panic(err)
}

/*
Syntax tells the user that they've made a syntax error
and cleanly exits the program afterwards.
*/
func Syntax(err SyntaxError) {
	fmt.Printf("Syntax error: %s\n", err.Desc)
	os.Exit(1)
}

/*
Connect informs the user about an error that has occured while
attempting to connect to the server.
*/
func Connect(err error) {
	fmt.Printf("Connection error: %s\n", err.Error())
	os.Exit(1)
}

/*
General is used for general errors.
*/
func General(err string) {
	fmt.Printf("General error: %s\n", err)
	os.Exit(1)
}

/*
Identity is used for Identity-related errors.
*/
func Identity(err error) {
	fmt.Printf("Identity error: %s\n", err)
	os.Exit(1)
}
