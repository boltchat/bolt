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
