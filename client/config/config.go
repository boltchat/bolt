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

package config

import (
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

var configFile = &File{
	Filename: "config.yml",
	Default:  *GetDefaultConfig(),
}

func parseConfig(raw []byte) (*Config, error) {
	config := &Config{}
	err := yaml.Unmarshal(raw, config)

	if err != nil {
		return nil, err
	}

	return config, nil
}

func LoadConfig() {
	configRaw, readErr := configFile.Read()
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
