// bolt.chat
// Copyright (C) 2021  Kees van Voorthuizen
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package client

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/bolt-chat/client/config"
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
			panic(parseErr)
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
