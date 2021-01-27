// bolt.chat
// Copyright (C) 2021  The bolt.chat Authors
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
	"path"

	"gopkg.in/yaml.v2"
)

type Identity struct {
	Nickname string `yaml:"nickname"`
}

type IdentityList map[string]Identity

var identityList IdentityList

const DefaultIdentity string = "default"

func getIdentityLocation() string {
	return path.Join(getConfigRoot(), "identity.yml")
}

func parseIdentityList(raw []byte) *IdentityList {
	identityList := &IdentityList{}
	err := yaml.Unmarshal(raw, identityList)

	if err != nil {
		panic(err)
	}

	return identityList
}

func readIdentityList() ([]byte, error) {
	raw, err := ioutil.ReadFile(getIdentityLocation())

	if err != nil {
		return nil, err
	}

	return raw, nil
}

func LoadIdentityList() {
	identityRaw, err := readIdentityList()

	if err != nil {
		panic(err)
	}

	identityList = *parseIdentityList(identityRaw)
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
