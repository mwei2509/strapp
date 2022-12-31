package orchestrator

import (
	"os"

	"github.com/mwei2509/strapp/pkg/apps/node"
)

type Database struct {
	Name string
	Type string // e.g. postgres mysql etc
	Port int64
}

type Service struct {
	Directory        string
	ContextDirectory string
	TemplateType     string
	IsRootApp        bool `yaml:"-"`
	Type             string
	Port             int64
	DebuggerPort     int64
	Name             string
	Language         string
	Framework        string `yaml:"framework,omitempty"`
	Css              string `yaml:"css,omitempty"`
	StateManagement  string `yaml:"state_management,omitempty"`
	Deployment       string `yaml:"deployment,omitempty"`
	Orm              string `yaml:"orm,omitempty"`
	Databases        []string
}

func (s *Service) init() error {
	if !s.IsRootApp {
		// create directories
		if err := s.createServiceDirectory(); err != nil {
			return err
		}
	}
	if err := s.createApp(); err != nil {
		return err
	}
	return nil
}

func (s *Service) createApp() error {
	switch {
	case s.Language == "typescript" || s.Language == "javascript":
		n := node.Node{
			Name:             s.Name,
			Port:             s.Port,
			DebuggerPort:     s.DebuggerPort,
			Language:         s.Language,
			Directory:        s.Directory,
			ContextDirectory: s.ContextDirectory,
			Framework:        s.Framework,
			Databases:        s.Databases,
		}
		if err := n.CreateNodeApp(); err != nil {
			return err
		}
	}
	return nil
}

func (s *Service) createServiceDirectory() error {
	err := os.Mkdir(s.Directory, 0750)
	if err != nil && !os.IsExist(err) {
		return err
	}
	return nil
}
