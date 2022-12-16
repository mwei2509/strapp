package aws

import (
	"github.com/mwei2509/strapp/pkg/utility"
)

var colorPurple string = "\033[35m"

func getPrefix() utility.PkgPrefix {
	return utility.PkgPrefix{
		Color: colorPurple,
		Name:  "[aws]",
	}
}
func Log(args ...any) {
	utility.Log(getPrefix(), args...)
}

func Fatal(args ...any) {
	utility.Fatal(getPrefix(), args...)
}

func Warn(args ...any) {
	utility.Warn(getPrefix(), args...)
}
