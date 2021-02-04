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
	"os"
	"strings"

	"github.com/bolt-chat/client/cli/cmd"
	"github.com/bolt-chat/client/cli/cmd/connect"
	"github.com/bolt-chat/client/cli/cmd/help"
	"github.com/bolt-chat/client/cli/cmd/identity"
	"github.com/bolt-chat/client/cli/cmd/version"
	"github.com/fatih/color"
)

var ErrTooFewArgs = errors.New("too few arguments")
var ErrCmdNotFound = errors.New("command not found")
var ErrSubCmdNotFound = errors.New("subcommand not found")

var commands = []*cmd.Command{
	help.HelpCommand,
	version.VersionCommand,
	connect.ConnectCommand,
	identity.IdentityCommand,
}

func formatCmd(cmd *cmd.Command) string {
	return fmt.Sprintf("%s %s\t%s", cmd.Name, cmd.Usage, cmd.Desc)
}

func PrintUsage() {
	fmt.Printf("usage: boltchat <command> [subcommand] [args...]\ncommands:\n")

	for _, cmd := range commands {
		// Print command details
		fmt.Printf("\t%s\n", formatCmd(cmd))

		if len(cmd.Subcommands) > 0 {
			for _, subcmd := range cmd.Subcommands {
				// Print subcommand details
				fmt.Printf("\t%s %s\n", cmd.Name, formatCmd(subcmd))
			}
		} else {
			// Print blank line after each group of subcommands
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
	// No command was given
	if len(args) == 0 {
		return nil, ErrTooFewArgs
	}

	// Get the command from the first argument
	cmd := getCmd(commands, args[0])
	if cmd == nil {
		return nil, ErrCmdNotFound
	}

	/*
		Print usage when issuing the 'help' command.
		This has to be handled here, because the
		`commands` array can not reference itself.
	*/
	if cmd == help.HelpCommand {
		cmd.Handler = func(args []string) error {
			PrintUsage()
			return nil
		}
		return cmd, nil
	}

	// Return the command itself if it doesn't
	// have any subcommands
	if len(cmd.Subcommands) == 0 {
		cmd.Args = args[1:]
		return cmd, nil
	}

	// No subcommand was given
	if len(args) < 2 {
		return nil, ErrTooFewArgs
	}

	// Get the subcommand from the second argument
	subcmd := getCmd(cmd.Subcommands, args[1])
	if subcmd == nil {
		return nil, ErrSubCmdNotFound
	}

	subcmd.Args = args[2:]
	return subcmd, nil
}

func HandleCommandError(cmdErr error) {
	fmt.Printf(color.RedString("Command error: %s\n\n", cmdErr))
	PrintUsage()
	os.Exit(1)
}
