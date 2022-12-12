package cmd

import (
	"github.com/mwei2509/strapp/pkg"
	"github.com/mwei2509/strapp/pkg/app"
	"github.com/spf13/cobra"
)

var flags app.Flag
var languageFlag string

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
	Run: func(cmd *cobra.Command, args []string) {
		if err := pkg.CreateApp(args[0], flags); err != nil {
			Fatal(err)
			return
		}
		Log("Success!")
	},
}

func init() {
	rootCmd.AddCommand(appCreateCommand)
	// flags
	appCreateCommand.Flags().StringArrayVarP(&flags.Type, "type", "t", app.FlagDefaults.Type, "type of app, e.g. api, web.  Use multiple for monorepo setup")
	appCreateCommand.Flags().StringVarP(&flags.Language, "language", "l", app.FlagDefaults.Language, "language")
	appCreateCommand.Flags().StringVarP(&flags.Framework, "framework", "f", app.FlagDefaults.Framework, "framework")
}
