package orchestrator

/**

App is the orchestrator of the application

**/
import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/mwei2509/strapp/pkg/ops"
	"golang.org/x/sync/errgroup"
)

type Orchestrator struct {
	Name      string
	Directory string
	Flags     Flag
	Config    Config
}

func (o *Orchestrator) init() error {
	// create in memory "strapprc" from flags
	// if strapprc already exists, ask permission to overwrite
	// > if Y - overwrite strapprc
	// > if N - ask to read from strapprc
	// > > if Y - overwrite strapprc
	// > > if N - exit
	// read config from strapRC

	// if strapprc exists, read from it

	// validations / config
	if _, err := os.Stat(o.Directory + "/.strapprc"); os.IsNotExist(err) {
		o.Config.Services = make([]Service, 0)
		for i, v := range o.Flags.Type {
			o.Config.Services = append(o.Config.Services, Service{
				Type:      v,
				Language:  o.Flags.Language[i],
				Framework: o.Flags.Framework[i],
			})
		}
	}
	return nil
}

type Flag struct {
	Type      []string
	Language  []string
	Framework []string
	Orm       []string
	Port      []int64
	Database  string
}

var FlagDefaults Flag = Flag{
	Type:      []string{"api"},
	Language:  []string{"typescript"},
	Framework: []string{"koa"},
	Orm:       []string{},
	Port:      []int64{},
}

func Do(directory string, flags Flag) error {
	// install necessary dependencies (return to this)
	installs := ops.GetInstallNeeded()
	if len(installs) > 0 {
		ops.InstallDependencies()
	}

	// init the app orchestrator
	var name string
	if directory != "." {
		nameSlice := strings.Split(directory, "/")
		name = nameSlice[len(nameSlice)-1]
	} else {
		path, err := os.Getwd()
		if err != nil {
			return err
		}
		nameSlice := strings.Split(path, "/")
		name = nameSlice[len(nameSlice)-1]
	}
	o := Orchestrator{Name: name, Directory: directory, Flags: flags}
	if err := o.init(); err != nil {
		return err
	}

	// create project directory
	if err := o.createAppDirectory(); err != nil {
		return err
	}

	// set app configs
	if err := o.setConfig(); err != nil {
		return err
	}

	// start docker compose
	o.Log("setting up docker...")
	if err := o.Config.DockerCompose.WriteDockerCompose(o.Directory); err != nil {
		return err
	}

	// init services
	eg := new(errgroup.Group)
	for i := 0; i < len(o.Config.Services); i++ {
		service := o.Config.Services[i]
		eg.Go(func() error {
			return service.init()
		})
	}
	if err := eg.Wait(); err != nil {
		return err
	}

	// init docker

	return nil
}

func (o Orchestrator) createAppDirectory() error {
	if o.Directory != "." {
		// return an error if app exists
		if _, err := os.Stat(o.Directory); !os.IsNotExist(err) {
			return errors.New(fmt.Sprintf("%s already exists", o.Directory))
		}

		err := os.Mkdir(o.Directory, 0750)
		if err != nil && !os.IsExist(err) {
			return err
		}
	}
	return nil
}
