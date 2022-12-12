package pkg

import "github.com/mwei2509/strapp/pkg/app"

func CreateApp(directory string, flags app.Flag) error {
	err := app.Do(directory, flags)
	if err != nil {
		return err
	}
	return nil
}
