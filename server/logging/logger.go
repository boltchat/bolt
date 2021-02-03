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

package logging

import (
	"fmt"
	"os"
	"time"

	"github.com/fatih/color"
)

type EventType int

const (
	RecvType EventType = iota
	SendType EventType = iota
)

func logBase(
	level string,
	msg string,
) {
	fmt.Printf("%s %s %s\n", color.HiBlackString(time.Now().Format("15:04:05")), level, msg)
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

func LogEvent(evtType EventType, evt interface{}) {
	typeMap := map[EventType]string{
		RecvType: color.HiCyanString("<--"),
		SendType: color.HiRedString("-->"),
	}

	logBase(color.HiMagentaString("EVENT"), fmt.Sprintf(
		"%s %v",
		typeMap[evtType],
		evt,
	))
}

// LogListener TODO
func LogListener(addr string) {
	LogInfo(fmt.Sprintf("Server listening on %s.", addr))
}
