package aws

import (
	"os"
	"strings"

	u "github.com/mwei2509/strapp/pkg/utility"
)

// https://docs.aws.amazon.com/lambda/latest/dg/images-test.html
func LambdaBuildSandbox(port string) error {
	// assume we're logged into aws already
	// assume aws is installed

	// get current directory
	path, err := os.Getwd()
	if err != nil {
		return err
	}
	pathArr := strings.Split(path, "/")
	dir := pathArr[len(pathArr)-1]

	if err := u.RunCommandLogOutput("docker build -t " + dir + " ."); err != nil {
		return err
	}

	log.Log("finished building, starting up sandbox on port " + port)
	log.Log("run POST requests to http://localhost:" + port + "/2015-03-31/functions/function/invocations")
	if err := u.RunCommandLogOutput("docker run --rm -p " + port + ":" + port + " projections"); err != nil {
		return err
	}
	return nil
}
