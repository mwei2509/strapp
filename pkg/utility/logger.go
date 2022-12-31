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

type Logger struct {
	pkgPrefix PkgPrefix
}

var log = Logger{
	pkgPrefix: PkgPrefix{
		Name:  "[util]",
		Color: "\033[35m",
	},
}

func CreatePackageLog(pkgPrefix PkgPrefix) Logger {
	return Logger{
		pkgPrefix: pkgPrefix,
	}
}

func (l Logger) Log(args ...any) {
	fmt.Println(l.wrapLog(Info, args...)...)
}

func (l Logger) Fatal(args ...any) {
	fmt.Println(l.wrapLog(Error, args...)...)
}

func (l Logger) Warn(args ...any) {
	fmt.Println(l.wrapLog(Warning, args...)...)
}

func (l Logger) wrapLog(t LogType, args ...any) []any {
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
	logs := append([]any{string(colorPrefix), "[[ STRAPP ]]", string(l.pkgPrefix.Color), l.pkgPrefix.Name, string(colorReset)}, args...)
	return logs
}
