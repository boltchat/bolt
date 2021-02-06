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

package prompt

import (
	"github.com/boltchat/lib/client"
	"github.com/gdamore/tcell/v2"
)

type Mode int

type ModeHandler func(s tcell.Screen, c *client.Client, evt tcell.Event)

const (
	MessageMode Mode = iota
)

var modeHandlers = map[Mode]ModeHandler{
	MessageMode: handleMessageMode,
}

var modeStrs = map[Mode]string{
	MessageMode: "Msg",
}
