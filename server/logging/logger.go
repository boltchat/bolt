package logging

import (
	"fmt"
	"time"

	"github.com/fatih/color"
)

func logBase(
	level string,
	msg string,
) {
	fmt.Printf("%s [%s] %s\n", color.HiBlackString(time.Now().Format("15:04:05")), level, msg)
}

func logInfo(msg string) {
	logBase(color.CyanString("INFO"), msg)
}

// LogListener TODO
func LogListener(ip string, port int) {
	logInfo(fmt.Sprintf("Server listening on %s:%v.", ip, port))
}
