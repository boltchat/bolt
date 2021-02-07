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

package markdown

import (
	"regexp"
	"strings"

	"github.com/fatih/color"
)

const BoldRegex string = "(\\*+)(\\s*\\b)([^\\*]*)(\\b\\s*)(\\*+)"

func BoldReplacer(r *regexp.Regexp) func(s string) string {
	return func(s string) string {
		s = strings.ReplaceAll(s, "*", "")
		return color.New(color.Bold).Sprintf(s)
	}
}
