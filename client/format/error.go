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

package format

import (
	"encoding/json"
	"fmt"

	"github.com/bolt-chat/protocol/errs"
	"github.com/bolt-chat/protocol/events"

	"github.com/fatih/color"
)

var errorMap = map[string]string{
	errs.InvalidEvent:    "This event type does not exist.",
	errs.InvalidFormat:   "The format of your request could not be parsed.",
	errs.TooManyMessages: "You're sending too many messages. Please slow down.",
	errs.Unidentified:    "You need to identify yourself before you can interact with this server.",
}

func FormatError(e *events.BaseEvent) string {
	errEvt := &events.ErrorEvent{}
	json.Unmarshal(*e.Raw, errEvt)

	err := errEvt.Error

	// A formatter exists for this error
	if format, ok := errorMap[errEvt.Error]; ok {
		err = format
	}

	return color.HiRedString(
		fmt.Sprintf("[!] %s", err),
	)
}
