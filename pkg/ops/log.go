package ops

import (
	"github.com/mwei2509/strapp/pkg/utility"
)

var colorGreen string = "\033[1;32m"
var appPrefix utility.PkgPrefix = utility.PkgPrefix{Name: "[ops]", Color: colorGreen}
var log = utility.CreatePackageLog(appPrefix)
