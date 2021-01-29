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

func getIdentityLocation() string {
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
	raw, err := ioutil.ReadFile(getIdentityLocation())

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

		writeErr := ioutil.WriteFile(getIdentityLocation(), defaultConf, 0644)
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
