package node

import (
	"github.com/mwei2509/strapp/pkg/utility"
)

var colorMagenta string = "\033[1;35m"
var nodePrefix utility.PkgPrefix = utility.PkgPrefix{Name: "[node]", Color: colorMagenta}

func Log(args ...any) {
	utility.Log(nodePrefix, args...)
}

func Fatal(args ...any) {
	utility.Fatal(nodePrefix, args...)
}

func Warn(args ...any) {
	utility.Warn(nodePrefix, args...)
}
