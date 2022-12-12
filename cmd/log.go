package cmd

import (
	"github.com/mwei2509/strapp/pkg/utility"
)

var colorYellow string = "\033[33m"
var cmdPrefix utility.PkgPrefix = utility.PkgPrefix{Name: "[cmd]", Color: colorYellow}

func Log(args ...any) {
	utility.Log(cmdPrefix, args...)
}

func Fatal(args ...any) {
	utility.Fatal(cmdPrefix, args...)
}

func Warn(args ...any) {
	utility.Warn(cmdPrefix, args...)
}
