package cmd

import (
	app "github.com/mwei2509/strapp/pkg/app"
	"github.com/spf13/cobra"
)

var appTypeFlags []string

// appCreateCommand represents the app:create command
var appCreateCommand = &cobra.Command{
	Use:   "app:create",
	Short: "Let's Boot(strapp) an App",
	Long: `
create an app given a directory:
	
Examples:
  strapp app:create my-app --type api
  strapp app:create my-app --type web
  strapp app:create my-app --type api --type web
`,
	Args: cobra.MatchAll(cobra.ExactArgs(1)),
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := app.Do(args[0], appTypeFlags); err != nil {
			return err
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(appCreateCommand)
	// flags
	appCreateCommand.Flags().StringArrayVarP(&appTypeFlags, "type", "-t", []string{}, "testing")
}
