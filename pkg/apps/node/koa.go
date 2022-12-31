package node

import (
	"github.com/mwei2509/strapp/pkg/ops"
	"github.com/mwei2509/strapp/pkg/templates"
)

func (n *Node) setupKoa() error {
	log.Log("loading templates...")
	templatesSlice, err := templates.LoadTemplates("koa/api_base")
	if err != nil {
		return err
	}
	for path, tmpl := range templatesSlice {
		if err := templates.CreateFiles(tmpl, n.Directory+"/"+path, n); err != nil {
			return err
		}
	}

	log.Log("installing...")
	if err := ops.NpmInstall(n.Directory); err != nil {
		return err
	}

	return nil
}
