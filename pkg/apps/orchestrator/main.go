package orchestrator

/**

App is the orchestrator of the application

**/
import (
	"os"

	ops "github.com/mwei2509/strapp/pkg/ops"
	u "github.com/mwei2509/strapp/pkg/utility"
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

func (o *Orchestrator) setServicesFromFlags() error {
	if _, err := os.Stat(o.Directory + "/.strapprc"); os.IsNotExist(err) {
		// check template first
		if o.Flags.TemplateType == "" && len(o.Flags.Type) == 0 {
			// there are no flags
			if len(o.Flags.Language) > 0 || len(o.Flags.Framework) > 0 || len(o.Flags.Orm) > 0 || len(o.Flags.Port) > 0 {
			}
		}
		if o.Flags.TemplateType != "" {
			// there is a template type
			// depending on a type, add necessary flags
		} else {
			// template type not included
		}
		// if everything is empty
		if len(o.Flags.Type) == 0 {
			// set defaults
			log.Log("hi")
		}
		//
	}
	return nil
}

type Flag struct {
	Type         []string
	Language     []string
	Framework    []string
	Orm          []string
	Port         []int64
	Database     []string
	TemplateType string
}

var FlagDefaults Flag = Flag{
	Type:         []string{"api"},
	Language:     []string{"typescript"},
	Framework:    []string{"koa"},
	Orm:          []string{},
	Port:         []int64{},
	Database:     []string{"postgres"},
	TemplateType: "create-users-with-login",
}

func Do(directory string, flags Flag) error {
	// install necessary dependencies (return to this)
	installs := ops.GetInstallNeeded()
	if len(installs) > 0 {
		ops.InstallDependencies()
	}

	var err error
	defer func() error {
		if err != nil {
			// cleanup
		}
		return err
	}()

	// init the app orchestrator
	name := u.GetNameFromDirectory(directory)
	o := Orchestrator{Name: name, Directory: directory, Flags: flags}
	// set flag defaults
	if err := o.setServicesFromFlags(); err != nil {
		return err
	}

	// probably not necessary
	// if err := o.init(); err != nil {
	// 	return err
	// }

	// create project directory
	if err = u.CreateDirectory(o.Directory); err != nil {
		return err
	}

	// set app configs
	if err = o.setConfig(); err != nil {
		return err
	}

	// YES OR NO FROM USER TO PROCEED
	// INCLUDE defer cleanup script?
	prompt := u.PromptContent{
		ErrMsg: "this is invalid",
		Label:  "this is a label",
	}
	input, err := prompt.GetInput()
	if err != nil {
		return err
	}
	log.Log(input)

	// start docker compose
	log.Log("setting up docker...")
	if err = o.Config.DockerCompose.WriteDockerCompose(o.Directory); err != nil {
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
	if err = eg.Wait(); err != nil {
		return err
	}

	// init docker

	return nil
}
