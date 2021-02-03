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
	"fmt"

	"github.com/bolt-chat/client/cli/cmd"
	"github.com/bolt-chat/client/cli/connect"
)

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
