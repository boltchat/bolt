// boltchat
// Copyright (C) 2021  The boltchat Authors
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

package config

import (
	"io/ioutil"
	"os"
	"path"

	"github.com/bolt-chat/client/errs"
	"gopkg.in/yaml.v2"
)

type Prompt struct {
	HOffset int `yaml:"hOffset"`
}

type Config struct {
	Prompt Prompt `yaml:"prompt"`
}

var config Config

func getConfigLocation() string {
	return path.Join(GetConfigRoot(), "config.yml")
}

func parseConfig(raw []byte) (*Config, error) {
	config := &Config{}
	err := yaml.Unmarshal(raw, config)

	if err != nil {
		return nil, err
	}

	return config, nil
}

// TODO: fix duplication
func readConfig() ([]byte, error) {
	configLocation := getConfigLocation()
	configRaw, err := ioutil.ReadFile(configLocation)

	if err != nil && !os.IsNotExist(err) {
		return nil, err
	}

	if len(configRaw) == 0 {
		configRoot := GetConfigRoot()
		defaultConf, marshalErr := yaml.Marshal(*GetDefaultConfig())

		if marshalErr != nil {
			return nil, marshalErr
		}

		stat, statErr := os.Stat(configRoot)
		if statErr != nil || !stat.IsDir() {
			os.MkdirAll(configRoot, 0755)
		}

		writeErr := ioutil.WriteFile(configLocation, defaultConf, 0644)
		if writeErr != nil {
			return nil, writeErr
		}

		configRaw = defaultConf
	}

	return configRaw, nil
}

func LoadConfig() {
	configRaw, readErr := readConfig()
	if readErr != nil {
		errs.Emerg(readErr)
	}

	conf, parseErr := parseConfig(configRaw)
	if parseErr != nil {
		errs.Emerg(parseErr)
	}

	config = *conf
}

func GetConfig() *Config {
	return &config
}
