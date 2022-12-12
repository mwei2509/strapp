package node

import (
	"fmt"
	"os/exec"

	"github.com/mwei2509/strapp/pkg/templates"
)

type Node struct {
	Directory string
	Framework string
}

func (n *Node) npmInit() error {
	cmd := exec.Command("npm", "init", "-y")
	cmd.Dir = n.Directory
	out, err := cmd.Output()
	if err != nil {
		return err
	}
	Log(fmt.Sprintf("%s\n", out))
	return nil
}

func (n *Node) createFramework() error {
	Log(n.Framework)
	return nil
}

func CreateNodeApp(directory string, framework string) (Node, error) {
	n := Node{Directory: directory, Framework: framework}

	// init
	if err := n.npmInit(); err != nil {
		return Node{}, err
	}

	// framework
	if err := n.createFramework(); err != nil {
		return Node{}, err
	}

	appjs := templates.AppJs{Port: 3000}
	err := templates.CreateVanillaAppJs(n.Directory, appjs)
	if err != nil {
		return Node{}, err
	}
	return n, nil
}
