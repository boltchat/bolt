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

package identity

import (
	"fmt"
	"strings"

	"github.com/boltchat/client/config"
	"github.com/boltchat/client/identity"
	"github.com/fatih/color"
)

// CreateIdentity creates a new Identity.
func CreateIdentity(identityID string) (*identity.Identity, error) {
	nickname := ""

	for strings.TrimSpace(identityID) == "" {
		fmt.Printf(
			color.HiCyanString("The Identity ID is used for referencing this Identity later on. An example would be %s or %s.\n"),
			color.HiYellowString("my_alt"),
			color.HiYellowString("very_very_secret"),
		)
		fmt.Printf("Identity ID: ")
		fmt.Scanln(&identityID)
	}

	for strings.TrimSpace(nickname) == "" {
		fmt.Printf("Nickname: ")
		fmt.Scanln(&nickname)
	}

	return identity.CreateIdentity(&config.Identity{
		Nickname: nickname,
	}, identityID)
}

// AskCreate will prompt the user if they'd like to create
// the missing identity.
func AskCreate(identityID string) bool {
	fmt.Printf(
		"Identity '%s' does not exist.\nWould you like to create it now? [Y/n] ",
		identityID,
	)

	answer := ""
	fmt.Scanln(&answer)

	return strings.ToLower(answer) == "y" || answer == ""
}
