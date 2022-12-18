package orchestrator

import (
	"github.com/mwei2509/strapp/pkg/utility"
)

var colorPurple string = "\033[35m"

func getPrefix() utility.PkgPrefix {
	return utility.PkgPrefix{
		Color: colorPurple,
		Name:  "[app]",
	}
}
func (o *Orchestrator) Log(args ...any) {

	utility.Log(getPrefix(), args...)
}

func (o *Orchestrator) Fatal(args ...any) {
	utility.Fatal(getPrefix(), args...)
}

func (o *Orchestrator) Warn(args ...any) {
	utility.Warn(getPrefix(), args...)
}
