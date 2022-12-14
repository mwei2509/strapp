package orchestrator

import (
	"os"

	"github.com/mwei2509/strapp/pkg/apps/runtime/node"
)

type Service struct {
	Directory       string
	IsRootApp       bool
	Type            string
	Port            int64
	Name            string
	Language        string
	Framework       string
	Css             string
	StateManagement string
	Deployment      string
	Orm             string
	Datastores      []string
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
			Name:      s.Name,
			Port:      s.Port,
			Language:  s.Language,
			Directory: s.Directory,
			Framework: s.Framework,
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
