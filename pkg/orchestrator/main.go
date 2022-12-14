package orchestrator

/**

App is the orchestrator of the application

**/
import (
	"errors"
	"fmt"
	"os"

	"github.com/mwei2509/strapp/pkg/ops"
	"golang.org/x/sync/errgroup"
)

type Orchestrator struct {
	Directory string
	Flags     Flag
	Config    Config
}

func (o *Orchestrator) init() error {
	o.Flags.SetDefaults()
	// validations / config
	return nil
}

type Flag struct {
	Type      []string
	Language  string
	Framework string
	Wonky     string
}

func (f *Flag) SetDefaults() {
	if f.Wonky == "" {
		f.Wonky = "i am wonky"
	}
}

var FlagDefaults Flag = Flag{
	Type:      []string{"api"},
	Language:  "typescript",
	Framework: "express",
}

func Do(directory string, flags Flag) error {
	// install necessary dependencies (return to this)
	installs := ops.GetInstallNeeded()
	if len(installs) > 0 {
		ops.InstallDependencies()
	}

	// init the app orchestrator
	o := Orchestrator{Directory: directory, Flags: flags}
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
