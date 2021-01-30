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

package util

import (
	"fmt"
	"strings"
)

// PrintBanner prints a neat ASCII art banner to stdout
func PrintBanner() {
	ascii := strings.Join([]string{
		" _           _ _              _           _   ",
		"| |         | | |            | |         | |  ",
		"| |__   ___ | | |_        ___| |__   __ _| |_ ",
		"| '_ \\ / _ \\| | __|      / __| '_ \\ / _` | __|",
		"| |_) | (_) | | |_   _  | (__| | | | (_| | |_ ",
		"|_.__/ \\___/|_|\\__| (_)  \\___|_| |_|\\__,_|\\__|",
	}, "\n")

	// Format & print the banner
	fmt.Printf("\n%s\n\n\n", ascii)
}
