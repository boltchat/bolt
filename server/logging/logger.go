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

package logging

import (
	"fmt"
	"os"
	"time"

	"github.com/fatih/color"
)

func logBase(
	level string,
	msg string,
) {
	fmt.Printf("%s [%s] %s\n", color.HiBlackString(time.Now().Format("15:04:05")), level, msg)
}

func LogInfo(msg string) {
	logBase(color.CyanString("INFO"), msg)
}

func LogError(msg string) {
	logBase(color.HiRedString("ERROR"), msg)
}

func LogDebug(msg string, data interface{}) {
	_, isDebug := os.LookupEnv("DEBUG")
	if !isDebug {
		return
	}

	if data != nil {
		msg = fmt.Sprintf("%s %v", msg, data)
	}

	logBase(color.HiYellowString("DEBUG"), msg)
}

func LogEvent(evt interface{}) {
	LogDebug("event:", evt)
}

// LogListener TODO
func LogListener(addr string) {
	LogInfo(fmt.Sprintf("Server listening on %s.", addr))
}
