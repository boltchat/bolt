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
	"errors"
	"io/ioutil"
	"os"
	"path"

	"github.com/bolt-chat/client/errs"
	"gopkg.in/yaml.v2"
)

type Identity struct {
	Nickname string `yaml:"nickname"`
}

type IdentityList map[string]Identity

var identityList IdentityList

const DefaultIdentity string = "default"

func GetIdentityLocation() string {
	return path.Join(GetConfigRoot(), "identity.yml")
}

func parseIdentityList(raw []byte) (*IdentityList, error) {
	identityList := &IdentityList{}
	err := yaml.Unmarshal(raw, identityList)

	if err != nil {
		return nil, err
	}

	return identityList, nil
}

// TODO: fix duplication
func readIdentityList() ([]byte, error) {
	raw, err := ioutil.ReadFile(GetIdentityLocation())

	if err != nil && !os.IsNotExist(err) {
		return nil, err
	}

	if len(raw) == 0 {
		configRoot := GetConfigRoot()
		defaultConf, marshalErr := yaml.Marshal(&IdentityList{
			DefaultIdentity: Identity{},
		})

		if marshalErr != nil {
			return nil, marshalErr
		}

		stat, statErr := os.Stat(configRoot)
		if statErr != nil || !stat.IsDir() {
			os.MkdirAll(configRoot, 0755)
		}

		writeErr := ioutil.WriteFile(GetIdentityLocation(), defaultConf, 0644)
		if writeErr != nil {
			return nil, writeErr
		}

		raw = defaultConf
	}

	return raw, nil
}

func LoadIdentityList() {
	identityRaw, readErr := readIdentityList()
	if readErr != nil {
		errs.Emerg(readErr)
	}

	identity, parseErr := parseIdentityList(identityRaw)
	if parseErr != nil {
		errs.Emerg(readErr)
	}

	identityList = *identity
}

func GetIdentityList() *IdentityList {
	return &identityList
}

func GetIdentity(id string) (*Identity, error) {
	if identity, ok := identityList[id]; ok {
		return &identity, nil
	}

	return nil, errors.New("identity not found")
}
