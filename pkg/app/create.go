package app

import (
	"errors"
	"fmt"
	"os"

	"github.com/mwei2509/strapp/pkg/node"
	"github.com/mwei2509/strapp/pkg/ops"
)

type App struct {
	Directory string
	Flags     Flag
	Config    Config
}

func (a *App) init(directory string, flags Flag) {
	// init app?
	a.Directory = directory
	a.Flags = flags
	a.Flags.SetDefaults()
	// validations / config
}

func (a *App) setConfig() error {
	config, err := getConf(a.Directory)
	if err != nil {
		return err
	}
	a.Config = config
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

	// init the app
	a := App{}
	a.init(directory, flags)

	// create project directory
	err := createAppDirectory(directory)
	if err != nil {
		return err
	}

	// set configs
	a.setConfig()

	// if single app => create single app
	// if monorepo =>

	// if node
	if _, err := node.CreateNodeApp(directory, flags.Framework); err != nil {
		return err
	}

	return nil
}

func createAppDirectory(directory string) error {
	if directory != "." {
		// return an error if app exists
		if _, err := os.Stat(directory); !os.IsNotExist(err) {
			return errors.New(fmt.Sprintf("%s already exists", directory))
		}

		err := os.Mkdir(directory, 0750)
		if err != nil && !os.IsExist(err) {
			return err
		}
	}
	return nil
}
