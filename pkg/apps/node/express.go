package node

import (
	"github.com/mwei2509/strapp/pkg/ops"
	"github.com/mwei2509/strapp/pkg/templates/files/express"
)

/* EXPRESS */

func (n *Node) setupExpress() error {
	if err := ops.NpmInstallPkg(n.Directory, "express"); err != nil {
		return err
	}

	appjs := express.AppJs{Port: n.Port}
	if err := express.CreateVanillaAppJs(n.Directory, appjs); err != nil {
		return err
	}
	return nil
}
