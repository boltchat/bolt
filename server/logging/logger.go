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

// LogListener TODO
func LogListener(ip string, port int) {
	LogInfo(fmt.Sprintf("Server listening on %s:%v.", ip, port))
}
