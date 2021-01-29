// bolt.chat
// Copyright (C) 2021  The bolt.chat Authors
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

package tui

import "github.com/gdamore/tcell/v2"

func splitChunks(str string, n int) []string {
	/*
		Return one chunk containing the entire string
		if the string does not exceed `n`.
	*/
	if len(str) < n {
		return []string{str}
	}

	chunks := make([]string, 0)

	for i := 0; i < len(str); i++ {
		if i%n == 0 {
			chunks = append(chunks, str[i:i*2])
		}
	}

	return chunks
}

func printLine(s tcell.Screen, y int, str string) {
	/*
		I do not like this workaround at all, but at this
		point, I've given up on trying to find a better
		solution. Feel free to create a Pull Request if
		you're able to improve this.
	*/
	chars := []rune("\b\b" + str)

	s.SetContent(0, y, ' ', chars[1:], tcell.StyleDefault)
}

func clearLine(s tcell.Screen, y int, w int) {
	// Clear every cell to `w`
	for x := 0; x < w; x++ {
		s.SetContent(x, y, ' ', nil, tcell.StyleDefault)
	}
}
