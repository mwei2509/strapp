package node

import (
	"fmt"
	"os/exec"
	// "github.com/mwei2509/strapp/pkg/templates"
)

/* REACT */
func (n *Node) installReact() error {
	cmd := exec.Command("npm", "install", "react")
	cmd.Dir = n.Directory
	out, err := cmd.Output()
	if err != nil {
		return err
	}
	n.Log(fmt.Sprintf("%s\n", out))
	return nil
}

func (n *Node) setupReact() error {
	n.Log("install Koa")
	// if err := n.installReact(); err != nil {
	// 	return err
	// }

	// appjs := templates.AppJs{Port: 3000}
	// if err := templates.CreateVanillaAppJs(n.Directory, appjs); err != nil {
	// 	return err
	// }
	return nil

}
