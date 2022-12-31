package utility

import (
	"fmt"
	"strings"

	"github.com/manifoldco/promptui"
)

type PromptContent struct {
	ErrMsg           string
	Label            string
	Prefix           PkgPrefix
	ValidInputs      []string
	CustomValidation func(input string) error
}

func CreatePrompt(setup PromptContent) func(input PromptContent) PromptContent {
	return func(input PromptContent) PromptContent {
		prompt := PromptContent{
			Prefix:           input.Prefix,
			ErrMsg:           input.ErrMsg,
			Label:            input.Label,
			ValidInputs:      input.ValidInputs,
			CustomValidation: input.CustomValidation,
		}
		if prompt.Prefix == (PkgPrefix{}) {
			prompt.Prefix = setup.Prefix
		}
		if prompt.ErrMsg == "" {
			prompt.ErrMsg = setup.ErrMsg
		}
		return prompt
	}
}

func formatLabel(pkgPrefix PkgPrefix, label string) []string {
	colorCyan := "\033[36m"
	colorYellow := "\033[33m"
	colorPurple := "\033[35m"
	colorLabel := colorYellow
	colorPrefix := colorCyan
	colorPrompt := colorPurple
	return []string{string(colorPrefix), "[[ STRAPP ]]", string(pkgPrefix.Color), pkgPrefix.Name, string(colorLabel), "[PROMPT] " + label, string(colorPrompt)}
}

func (pc *PromptContent) Validate(input string) error {
	if pc.CustomValidation != nil {
		return pc.CustomValidation(input)
	}

	return nil
}

func (pc *PromptContent) GetInput() (string, error) {
	templates := &promptui.PromptTemplates{
		Prompt:  "{{ . }} ",
		Valid:   "{{ . | green }} ",
		Invalid: "{{ . | red }} ",
		Success: "{{ . | bold }} ",
	}

	prompt := promptui.Prompt{
		Label:     strings.Join(formatLabel(pc.Prefix, pc.Label)[:], " "),
		Templates: templates,
		Validate:  pc.Validate,
	}

	result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return "", err
	}

	return result, nil
}

func (pc *PromptContent) Confirm() error {
	prompt := promptui.Prompt{
		Label:     strings.Join(formatLabel(pc.Prefix, pc.Label)[:], " "),
		IsConfirm: true,
	}

	_, err := prompt.Run()
	if err != nil {
		return err
	}
	return nil
}
