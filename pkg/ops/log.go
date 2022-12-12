package ops

import (
	"github.com/mwei2509/strapp/pkg/utility"
)

var colorGreen string = "\033[1;32m"
var appPrefix utility.PkgPrefix = utility.PkgPrefix{Name: "[ops]", Color: colorGreen}

func Log(args ...any) {
	utility.Log(appPrefix, args...)
}

func Fatal(args ...any) {
	utility.Fatal(appPrefix, args...)
}

func Warn(args ...any) {
	utility.Warn(appPrefix, args...)
}
