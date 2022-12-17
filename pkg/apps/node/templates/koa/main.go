package koa

import (
	"os"
	"text/template"
)

type KoaApp struct {
	Name                       string
	Directory                  string
	Port                       int64
	DebuggerPort               int64
	OrmPackageJsonScripts      string
	OrmPackageJsonDependencies string
	DbPackageJsonDependencies  string
}

func (k KoaApp) createFileFromTemplate(tmpl *template.Template, filepath string) error {
	file, err := os.Create(k.Directory + filepath)
	if err != nil {
		return err
	}

	err = tmpl.Execute(file, k)
	if err != nil {
		return err
	}
	return nil
}
