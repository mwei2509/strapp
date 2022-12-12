package utility

import (
	"fmt"
)

type LogType int

const (
	Info LogType = iota
	Error
	Warning
)

type PkgPrefix struct {
	Name  string
	Color string
}

func Log(pkgPrefix PkgPrefix, args ...any) {
	fmt.Println(WrapLog(Info, pkgPrefix, args...)...)
}

func Fatal(pkgPrefix PkgPrefix, args ...any) {
	fmt.Println(WrapLog(Error, pkgPrefix, args...)...)
}

func Warn(pkgPrefix PkgPrefix, args ...any) {
	fmt.Println(WrapLog(Warning, pkgPrefix, args...)...)
}

func WrapLog(t LogType, pkgPrefix PkgPrefix, args ...any) []any {
	colorWhite := "\033[0m"

	colorRed := "\033[31m"
	colorYellow := "\033[33m"
	colorBlue := "\033[34m"
	colorPurple := "\033[35m"
	colorCyan := "\033[36m"

	var colorPrefix string
	var colorReset string
	switch t {
	case Info:
		colorPrefix = colorCyan
		colorReset = colorWhite
	case Error:
		colorPrefix = colorRed
		colorReset = colorPurple
	case Warning:
		colorPrefix = colorYellow
		colorReset = colorBlue
	}
	logs := append([]any{string(colorPrefix), "[[ STRAPP ]]", string(pkgPrefix.Color), pkgPrefix.Name, string(colorReset)}, args...)
	return logs
}
