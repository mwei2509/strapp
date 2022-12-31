package cli

import (
	"github.com/mwei2509/strapp/pkg/utility"
	u "github.com/mwei2509/strapp/pkg/utility"
)

var colorPurple string = "\033[35m"
var pkgPrefix = u.PkgPrefix{
	Color: colorPurple,
	Name:  "[cli]",
}
var log = utility.CreatePackageLog(pkgPrefix)
