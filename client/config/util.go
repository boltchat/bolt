package config

import (
	"os"
	"path"
)

func getConfigRoot() string {
	root, err := os.UserConfigDir()
	if err != nil {
		panic(err)
	}

	return path.Join(root, "bolt.chat")
}
