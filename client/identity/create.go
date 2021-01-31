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

	"github.com/bolt-chat/client/config"
)

// CreateIdentity creates a new Identity.
func CreateIdentity(identityID string) *config.Identity {
	nickname := ""

	for strings.TrimSpace(nickname) == "" {
		fmt.Printf("Nickname: ")
		fmt.Scanln(&nickname)
	}

	identity := &config.Identity{
		Nickname: nickname,
	}

	identityList := *config.GetIdentityList()
	identityList[identityID] = *identity

	// Write changes to disk
	config.IdentityFile.Write(identityList)

	return identity
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