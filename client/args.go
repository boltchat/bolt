package client

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/keesvv/bolt.chat/client/config"
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
