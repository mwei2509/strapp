package templates

import (
	"embed"
	"text/template"

	"io/fs"
	"strings"
)

var (
	//go:embed files/*
	files     embed.FS
	templates map[string]*template.Template
)

// LoadTemplates("templates/koa/api_base")
// https://charly3pins.dev/blog/learn-how-to-use-the-embed-package-in-go-by-building-a-web-page-easily/
// embedded templates? https://levelup.gitconnected.com/using-go-templates-for-effective-web-development-f7df10b0e4a0
func LoadTemplates(root string) (map[string]*template.Template, error) {
	if templates == nil {
		templates = make(map[string]*template.Template)
	}

	if err := fs.WalkDir(files, "files/"+root, func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}

		pt, err := template.ParseFS(files, path)
		if err != nil {
			return err
		}
		path = strings.ReplaceAll(path, "files/"+root+"/", "")
		path = strings.ReplaceAll(path, ".tmpl", "")
		path = strings.Replace(path, "dot.", ".", 1)
		templates[path] = pt
		return nil
	}); err != nil {
		return nil, err
	}
	return templates, nil
}
