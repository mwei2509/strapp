package node

import (
	"os"
	"text/template"

	"github.com/mwei2509/strapp/pkg/apps/templates"
)

func (n *Node) setupKoa() error {
	templates, err := templates.LoadTemplates("koa/api_base")
	if err != nil {
		return err
	}
	for path, tmpl := range templates {
		if err := n.createFiles(tmpl, path); err != nil {
			return err
		}
	}

	if err := n.npmInstall(); err != nil {
		return err
	}

	return nil
}

func (n *Node) createFiles(tmpl *template.Template, path string) error {
	file, err := os.Create(n.Directory + "/" + path)
	if err != nil {
		return err
	}

	err = tmpl.Execute(file, n)
	if err != nil {
		return err
	}
	return nil
}
