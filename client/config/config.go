package config

import (
	"io/ioutil"
	"os"
	"path"

	"gopkg.in/yaml.v2"
)

type Config struct {
	A int `yaml:"a"`
	B int `yaml:"b"`
}

var config Config

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
	err := yaml.Unmarshal(raw, config)

	if err != nil {
		panic(err)
	}

	return config
}

func readConfig() ([]byte, error) {
	configRaw, err := ioutil.ReadFile(getConfigLocation())

	if err != nil {
		return nil, err
	}

	return configRaw, nil
}

func LoadConfig() {
	configRaw, _ := readConfig()

	if configRaw != nil {
		config = *parseConfig(configRaw)
	} else {
		config = Config{}
	}
}

func GetConfig() *Config {
	return &config
}
