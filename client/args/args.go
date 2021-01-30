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

package args

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/bolt-chat/client/config"
	"github.com/bolt-chat/client/errs"
	"github.com/fatih/color"
)

type Args struct {
	Hostname string
	Port     int
	Identity string
}

func printUsage() {
	fmt.Println("usage: boltchat <host> [identity]")
}

func GetArgs() *Args {
	rawArgs := os.Args[1:]

	// Set identity to 'default' by default
	identity := config.DefaultIdentity

	if len(rawArgs) < 1 {
		printUsage()
		os.Exit(1)
	} else if len(rawArgs) > 1 {
		identity = rawArgs[1]
	}

	splitHost := strings.Split(rawArgs[0], ":")
	hostname := splitHost[0]

	// The default port
	port := 3300

	// Custom port number is specified
	if len(splitHost) == 2 {
		parsedPort, parseErr := strconv.ParseInt(splitHost[1], 10, 32)

		if parseErr != nil {
			errs.Syntax(errs.SyntaxError{
				Error: parseErr,
				Desc: fmt.Sprintf(
					"You have entered an invalid host. Hosts "+
						"must be in the following format: %s.",
					color.HiCyanString("hostname|ip[:port]"),
				),
			})
		}

		port = int(parsedPort)
	}

	args := &Args{
		Hostname: hostname,
		Port:     port,
		Identity: identity,
	}

	return args
}
