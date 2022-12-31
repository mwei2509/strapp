package cli

import (
	i "github.com/mwei2509/strapp/pkg/ops"
	u "github.com/mwei2509/strapp/pkg/utility"
	"github.com/spf13/viper"
)

type CliOrchestrator struct {
	name       string
	directory  string
	language   string
	deployment string // brew vs npm
}

type CliFlags struct {
	Language   string
	Deployment string
}

var CliFlagDefaults CliFlags = CliFlags{
	Language:   "go",
	Deployment: "brew",
}

var createPrompt func(input u.PromptContent) u.PromptContent = u.CreatePrompt(u.PromptContent{
	Prefix: pkgPrefix,
	ErrMsg: "Invalid Response",
})

// var prompt u.PromptContent = u.PromptContent{ErrMsg: "Invalid Input", Prefix: getPrefix()}

func Do(directory string, flags CliFlags) error {
	name := u.GetNameFromDirectory(directory)
	// if go, then directory needs to be in gopath
	o := CliOrchestrator{name: name, directory: directory, language: flags.Language}

	switch o.language {
	case "go":
	}
	log.Log(o.directory)
	log.Log(name)

	viper.SetDefault("GoPath", "")
	viper.AutomaticEnv()
	test := viper.Get("GoPath")
	log.Log(test)

	// YES OR NO FROM USER TO PROCEED
	// INCLUDE defer cleanup script?
	prompt := createPrompt(u.PromptContent{Label: "select something"})
	// prompt.CustomValidation = func(input string) error {
	// 	if len(input) <= 0 {
	// 		return errors.New(prompt.ErrMsg)
	// 	}
	// 	return nil
	// }
	input, err := prompt.GetInput()
	if err != nil {
		return err
	}
	log.Log(input)

	prompt = createPrompt(u.PromptContent{Label: "Proceed?"})
	err = prompt.Confirm()
	if err != nil {
		return err
	}

	// login to gh
	i.LoginGithubApi()
	// u.CreateDirectory(o.directory)
	return nil
}
