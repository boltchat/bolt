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
	"github.com/boltchat/client/config"
	"github.com/gdamore/tcell/v2"
)

func SplitChunks(str string, n int) []string {
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
			end := i + n

			if end > len(str) {
				end = len(str)
			}

			chunks = append(chunks, str[i:end])
		}
	}

	return chunks
}

func PrintLine(s tcell.Screen, y int, str string) {
	/*
		I do not like this workaround at all, but at this
		point, I've given up on trying to find a better
		solution. Feel free to create a Pull Request if
		you're able to improve this.
	*/
	chars := []rune("\b\b" + str)

	s.SetContent(0, y, ' ', chars[1:], tcell.StyleDefault)
}

func ClearLine(s tcell.Screen, y int, w int) {
	// Clear every cell to `w`
	for x := 0; x < w; x++ {
		s.SetContent(x, y, ' ', nil, tcell.StyleDefault)
	}
}

func ClearBuffer(s tcell.Screen) {
	w, h := s.Size()
	hBuff := h - config.GetConfig().Prompt.HOffset

	// Clear the buffer
	for y := 0; y < hBuff; y++ {
		ClearLine(s, y, w)
	}
}
