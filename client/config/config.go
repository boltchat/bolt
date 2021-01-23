package config

import (
	"io/ioutil"
	"os"
	"path"

	"gopkg.in/yaml.v2"
)

type Prompt struct {
	HOffset int `yaml:"hOffset"`
}

type Config struct {
	Prompt Prompt `yaml:"prompt"`
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
	configLocation := getConfigLocation()
	configRaw, err := ioutil.ReadFile(configLocation)

	if err != nil && !os.IsNotExist(err) {
		return nil, err
	}

	if len(configRaw) == 0 {
		defaultConf, marshalErr := yaml.Marshal(*GetDefaultConfig())
		if marshalErr != nil {
			panic(marshalErr) // TODO
		}

		writeErr := ioutil.WriteFile(configLocation, defaultConf, 0644)
		if writeErr != nil {
			panic(writeErr) // TODO
		}

		configRaw = defaultConf
	}

	return configRaw, nil
}

func LoadConfig() {
	configRaw, _ := readConfig()

	if configRaw != nil {
		config = *parseConfig(configRaw)
	} else {
		config = *GetDefaultConfig()
	}
}

func GetConfig() *Config {
	return &config
}
