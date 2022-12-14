package node

import (
	"github.com/mwei2509/strapp/pkg/utility"
)

var colorMagenta string = "\033[1;35m"
var nodePrefix utility.PkgPrefix = utility.PkgPrefix{Color: colorMagenta}

func getPrefix(name string) utility.PkgPrefix {
	return utility.PkgPrefix{
		Color: colorMagenta,
		Name:  "[node::" + name + "]",
	}
}
func (n *Node) Log(args ...any) {

	utility.Log(getPrefix(n.Name), args...)
}

func (n *Node) Fatal(args ...any) {
	utility.Fatal(getPrefix(n.Name), args...)
}

func (n *Node) Warn(args ...any) {
	utility.Warn(getPrefix(n.Name), args...)
}
