package app

import (
	"errors"
	"fmt"
	"os"
)

type SingleApp struct {
	Type     string
	Language string
}
type Flag struct {
	Type     []string
	Language string
}

var FlagDefaults Flag = Flag{
	Type:     []string{"api"},
	Language: "typescript",
}

func Do(directory string, flags Flag) error {
	fmt.Println(wrapLog("create called", flags))
	fmt.Println(wrapLog(flags))

	// install necessary dependencies (return to this)
	installs := getInstallNeeded()
	if len(installs) > 0 {
		installDependencies()
	}

	// create project directory
	err := createAppDirectory(directory)
	if err != nil {
		return err
	}
	return nil
}

func createAppDirectory(directory string) error {
	// return an error if app exists
	if _, err := os.Stat(directory); !os.IsNotExist(err) {
		return errors.New(fmt.Sprintf("%s already exists", directory))
	}

	err := os.Mkdir(directory, 0750)
	if err != nil && !os.IsExist(err) {
		return err
	}
	return nil
}
