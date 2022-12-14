package node

import (
	"fmt"
	"os/exec"

	"github.com/mwei2509/strapp/pkg/apps/templates/node/express"
	// "github.com/mwei2509/strapp/pkg/templates"
)

/* EXPRESS */
func (n *Node) installExpress() error {
	cmd := exec.Command("npm", "install", "express")
	cmd.Dir = n.Directory
	out, err := cmd.Output()
	if err != nil {
		return err
	}
	n.Log(fmt.Sprintf("%s\n", out))
	return nil
}

func (n *Node) setupExpress() error {
	if err := n.installExpress(); err != nil {
		return err
	}

	appjs := express.AppJs{Port: n.Port}
	if err := express.CreateVanillaAppJs(n.Directory, appjs); err != nil {
		return err
	}
	return nil
}
