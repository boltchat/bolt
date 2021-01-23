package client

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Args struct {
	Hostname string
	Port     int
	Nickname string // TODO: this will change very soon
}

func printUsage() {
	fmt.Println("usage: boltchat <host> <nickname>")
}

func GetArgs() *Args {
	rawArgs := os.Args[1:]

	if len(rawArgs) < 2 {
		printUsage()
		os.Exit(1)
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
		Nickname: rawArgs[1],
	}

	return args
}
