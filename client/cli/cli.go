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

package cli

import (
	"errors"
	"fmt"
	"strings"

	"github.com/bolt-chat/client/cli/cmd"
	"github.com/bolt-chat/client/cli/connect"
)

var ErrTooFewArgs = errors.New("too few arguments")
var ErrCmdNotFound = errors.New("command not found")
var ErrSubCmdNotFound = errors.New("subcommand not found")

var commands = []*cmd.Command{
	connect.ConnectCommand,
}

func PrintUsage() {
	fmt.Printf("usage: boltchat <command> [subcommand] [args...]\ncommands:\n")

	for _, cmd := range commands {
		fmt.Println("\t", cmd.Name, "\t", cmd.Desc)

		if len(cmd.Subcommands) == 0 {
			fmt.Println()
		}
	}
}

func getCmd(cmds []*cmd.Command, rawCmd string) *cmd.Command {
	for _, cmd := range cmds {
		if cmd.Name == strings.ToLower(rawCmd) {
			return cmd
		}
	}

	return nil
}

func ParseCommand(args []string) (*cmd.Command, error) {
	if len(args) == 0 {
		return nil, ErrTooFewArgs
	}

	// Get the command from the first argument
	cmd := getCmd(commands, args[0])
	if cmd == nil {
		return nil, ErrCmdNotFound
	}

	if len(cmd.Subcommands) == 0 {
		cmd.Args = args[1:]
		return cmd, nil
	}

	// Get the subcommand from the second argument
	subcmd := getCmd(cmd.Subcommands, args[1])
	if subcmd == nil {
		return nil, ErrSubCmdNotFound
	}

	subcmd.Args = args[2:]
	return subcmd, nil
}
