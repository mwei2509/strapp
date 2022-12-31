package cmd

import (
	"github.com/mwei2509/strapp/pkg/utility"
)

var colorYellow string = "\033[33m"
var cmdPrefix utility.PkgPrefix = utility.PkgPrefix{Name: "[cmd]", Color: colorYellow}
var log = utility.CreatePackageLog(cmdPrefix)
