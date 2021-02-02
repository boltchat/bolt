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

	"github.com/bolt-chat/client/errs"
	"gopkg.in/yaml.v2"
)

var ErrNoSuchIdentity = errors.New("identity not found")

type Identity struct {
	Nickname   string `yaml:"nickname"`
	EntityPath string `yaml:"entity_path,omitempty"`
}

type IdentityList map[string]Identity

var identityList IdentityList

var IdentityFile = &File{
	Filename: "identity.yml",
	Default:  IdentityList{},
}

const DefaultIdentity string = "default"

func parseIdentityList(raw []byte) (*IdentityList, error) {
	identityList := &IdentityList{}
	err := yaml.Unmarshal(raw, identityList)

	if err != nil {
		return nil, err
	}

	return identityList, nil
}

func LoadIdentityList() {
	identityRaw, readErr := IdentityFile.Read()
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

func GetIdentityEntry(id string) (*Identity, error) {
	if identity, ok := identityList[id]; ok {
		return &identity, nil
	}

	return nil, ErrNoSuchIdentity
}
