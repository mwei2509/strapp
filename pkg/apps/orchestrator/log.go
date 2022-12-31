package orchestrator

import (
	"github.com/mwei2509/strapp/pkg/utility"
)

var colorPurple string = "\033[35m"

var log = utility.CreatePackageLog(utility.PkgPrefix{
	Color: colorPurple,
	Name:  "[app]",
})
