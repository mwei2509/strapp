package node

import (
	"github.com/mwei2509/strapp/pkg/apps/node/templates/koa"
)

func (n *Node) setupKoa() error {
	koaApp := koa.KoaApp{
		Name:         n.Name,
		Directory:    n.Directory,
		Port:         n.Port,
		DebuggerPort: n.Port + 2000,
	}
	if err := koaApp.CreateFooBarFiles(); err != nil {
		return err
	}

	if err := n.npmInstall(); err != nil {
		return err
	}

	return nil
}

// koa app
