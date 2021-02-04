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

package format

import (
	"fmt"

	"github.com/bolt-chat/protocol/errs"
	"github.com/bolt-chat/protocol/events"
	"github.com/mitchellh/mapstructure"

	"github.com/fatih/color"
)

var errorMap = map[string]string{
	errs.InvalidEvent:    "This event type does not exist.",
	errs.InvalidFormat:   "The format of your request could not be parsed.",
	errs.TooManyMessages: "You're sending too many messages. Please slow down.",
	errs.Unidentified:    "You need to identify yourself before you can interact with this server.",
}

func FormatError(e *events.Event) string {
	errData := events.ErrorData{}
	mapstructure.Decode(e.Data, &errData)

	err := errData.Error

	// A formatter exists for this error
	if format, ok := errorMap[errData.Error]; ok {
		err = format
	}

	return color.HiRedString(
		fmt.Sprintf("[!] %s", err),
	)
}
