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

package commands

import (
	"errors"

	"github.com/boltchat/server/pools"
)

type CommandHandler func(p *pools.ConnPool, c *pools.Connection, args []string)

var ErrCmdNotFound = errors.New("command not found")

var commands = map[string]CommandHandler{
	"ping": handlePing,
}

func Parse(cmd string) (CommandHandler, error) {
	cmdHandler, ok := commands[cmd]
	if !ok {
		return nil, ErrCmdNotFound
	}

	return cmdHandler, nil
}
