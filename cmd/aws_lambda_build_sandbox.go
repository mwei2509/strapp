package cmd

import (
	"github.com/mwei2509/strapp/pkg/aws"
	"github.com/spf13/cobra"
)

// awsLambdaBuildSandbox represents the app:create command
var awsLambdaBuildSandbox = &cobra.Command{
	Use:   "aws:lambda-build-sandbox",
	Short: "Let's Boot(strapp) an App",
	Long: `
run from the directory with your lambda dockerfile
	
Examples:
  strapp aws:lambda-build-sandbox --port 8080
`,
	Run: func(cmd *cobra.Command, args []string) {
		port, _ := cmd.Flags().GetString("port")

		if err := aws.LambdaBuildSandbox(port); err != nil {
			log.Fatal(err)
			return
		}
		log.Log("Success!")
	},
}

func init() {
	rootCmd.AddCommand(awsLambdaBuildSandbox)
	// flags
	awsLambdaBuildSandbox.Flags().StringP("port", "p", "8080", "port, defaults to 8080")
}
