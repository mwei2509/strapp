package cmd

import (
	cli "github.com/mwei2509/strapp/pkg/cli"
	"github.com/spf13/cobra"
)

var cliFlags cli.CliFlags

// cliCreateCommand represents the app:create command
var cliCreateCommand = &cobra.Command{
	Use:   "cli:create",
	Short: "Let's Boot(strapp) a CLI",
	Long: `
create a cli given a cli name:
	
Examples:
  strapp cli:create my-cli --language go --deployment brew
  strapp cli:create my-cli --language typescript --deployment npm
`,
	Args: cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		if err := cli.Do(args[0], cliFlags); err != nil {
			log.Fatal(err)
			return
		}
		log.Log("Success!")
	},
}

func init() {
	rootCmd.AddCommand(cliCreateCommand)
	// flags
	cliCreateCommand.Flags().StringVarP(&cliFlags.Language, "language", "l", cli.CliFlagDefaults.Language, "language")
	cliCreateCommand.Flags().StringVarP(&cliFlags.Deployment, "deployment", "d", cli.CliFlagDefaults.Deployment, "deployment")
	cliCreateCommand.Flags().Bool("dir-only", false, "only set up the directory, everything else will be done manually")
}
