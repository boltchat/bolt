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

package plugins

import (
	"errors"

	"github.com/boltchat/protocol/events"
	"github.com/boltchat/server/pools"
)

type NicknameValidationPlugin struct {
	MinChars int
	MaxChars int
}

func (NicknameValidationPlugin) GetInfo() *PluginInfo {
	return &PluginInfo{
		Id: "nickname-validation",
	}
}

var ErrNicknameTooShort = errors.New("nickname too short")
var ErrNicknameTooLong = errors.New("nickname too long")

func (p NicknameValidationPlugin) OnMessage(msg *events.MessageData, c *pools.Connection) error {
	return nil
}

func (p NicknameValidationPlugin) OnIdentify(data *events.JoinData, c *pools.Connection) error {
	if len(data.User.Nickname) < p.MinChars {
		return ErrNicknameTooShort
	}

	if len(data.User.Nickname) > p.MaxChars {
		return ErrNicknameTooLong
	}

	return nil
}
