package config

import (
	"os"
	"path"

	"gopkg.in/yaml.v2"
)

type Config struct {
	A int `yaml:"a"`
	B int `yaml:"b"`
}

var Current Config

func getConfigRoot() string {
	root, err := os.UserConfigDir()
	if err != nil {
		panic(err)
	}

	return path.Join(root, "bolt.chat")
}

func getConfigLocation() string {
	return path.Join(getConfigRoot(), "config.yml")
}

func parseConfig(raw []byte) *Config {
	config := &Config{}
	yaml.Unmarshal(raw, config)
	return config
}

func readConfig() ([]byte, error) {
	f, err := os.Open(getConfigLocation())
	configRaw := make([]byte, 1024)

	if err != nil {
		return nil, err
	}

	_, readErr := f.Read(configRaw)
	if readErr != nil {
		return nil, readErr
	}

	return configRaw, nil
}

func LoadConfig() {
	// TODO
	Current = *parseConfig([]byte("a: 1\nb: 2"))
}
