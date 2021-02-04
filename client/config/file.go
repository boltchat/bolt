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
	"io/ioutil"
	"os"
	"path"

	"github.com/boltchat/client/errs"
	"gopkg.in/yaml.v2"
)

// File is used to describe a YAML config file.
type File struct {
	Filename string
	Default  interface{}
}

// GetConfigRoot returns the base folder where all
// config files reside.
func GetConfigRoot() string {
	root, err := os.UserConfigDir()
	if err != nil {
		errs.Emerg(err)
	}

	return path.Join(root, "boltchat")
}

func (f *File) GetLocation() string {
	return path.Join(GetConfigRoot(), f.Filename)
}

func (f *File) Read() ([]byte, error) {
	configLocation := f.GetLocation()
	configRaw, err := ioutil.ReadFile(configLocation)

	if err != nil && !os.IsNotExist(err) {
		return nil, err
	}

	if len(configRaw) == 0 {
		writeRes, writeErr := f.Write(f.Default)
		if writeErr != nil {
			return nil, writeErr
		}

		configRaw = writeRes
	}

	return configRaw, nil
}

func (f *File) Write(data interface{}) ([]byte, error) {
	configRoot := GetConfigRoot()
	configLocation := f.GetLocation()
	conf, marshalErr := yaml.Marshal(data)

	if marshalErr != nil {
		return nil, marshalErr
	}

	stat, statErr := os.Stat(configRoot)
	if statErr != nil || !stat.IsDir() {
		os.MkdirAll(configRoot, 0755)
	}

	writeErr := ioutil.WriteFile(configLocation, conf, 0644)
	if writeErr != nil {
		return nil, writeErr
	}

	return conf, nil
}
