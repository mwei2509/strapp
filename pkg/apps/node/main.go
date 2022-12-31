package node

import (
	"os"
	"strings"

	"github.com/mwei2509/strapp/pkg/ops"
)

type Node struct {
	Name             string
	Port             int64
	DebuggerPort     int64
	ContextDirectory string
	Directory        string
	Language         string
	Framework        string
	Databases        []string
	Orm              string
}

func (n *Node) CreateNodeApp() error {
	if len(n.Databases) > 0 {
		// do sonething
	}

	// all node apps should have a src directory
	if n.Framework != "react" {
		err := os.MkdirAll(n.Directory+"/src", 0755)
		if err != nil && !os.IsExist(err) {
			return err
		}
	}

	// all node apps will be on node 18
	if err := ops.InstallNode(); err != nil {
		return err
	}

	// framework
	if err := n.createFramework(); err != nil {
		return err
	}

	// update docker compose with service stuff?

	return nil
}

func (n *Node) createFramework() error {
	switch strings.ToLower(n.Framework) {
	case "express":
		if err := n.setupExpress(); err != nil {
			return err
		}
	case "koa":
		if err := n.setupKoa(); err != nil {
			return err
		}
	case "react":
		if err := n.setupReact(); err != nil {
			return err
		}
	}
	return nil
}
