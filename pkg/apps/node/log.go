package node

import (
	"github.com/mwei2509/strapp/pkg/utility"
)

var colorMagenta string = "\033[1;35m"
var nodePrefix utility.PkgPrefix = utility.PkgPrefix{Color: colorMagenta}

var log = utility.CreatePackageLog(utility.PkgPrefix{
	Color: colorMagenta,
	Name:  "[node]",
})
