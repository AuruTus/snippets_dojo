package cfmt

import "fmt"

var (
	InfoStr = Teal
	WarnStr = Yellow
	FataStr = Red
)

var (
	Black   = ColorStr("\033[1;30m%s\033[0m")
	Red     = ColorStr("\033[1;31m%s\033[0m")
	Green   = ColorStr("\033[1;32m%s\033[0m")
	Yellow  = ColorStr("\033[1;33m%s\033[0m")
	Purple  = ColorStr("\033[1;34m%s\033[0m")
	Magenta = ColorStr("\033[1;35m%s\033[0m")
	Teal    = ColorStr("\033[1;36m%s\033[0m")
	White   = ColorStr("\033[1;37m%s\033[0m")
)

func ColorStr(colorString string) func(string, ...interface{}) string {
	return func(format string, args ...interface{}) string {
		return fmt.Sprintf(
			colorString,
			fmt.Sprintf(format, args...),
		)
	}
}
