package app

import (
	"github.com/mwei2509/strapp/pkg/utility"
)

var colorPurple string = "\033[35m"
var appPrefix utility.PkgPrefix = utility.PkgPrefix{Name: "[app]", Color: colorPurple}

func Log(args ...any) {
	utility.Log(appPrefix, args...)
}

func Fatal(args ...any) {
	utility.Fatal(appPrefix, args...)
}

func Warn(args ...any) {
	utility.Warn(appPrefix, args...)
}
