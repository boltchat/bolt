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
