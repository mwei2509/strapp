package aws

import (
	"bufio"
	"os"
	"os/exec"
	"strings"
)

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

	cmd := exec.Command("bash", "-c", "docker build -t "+dir+" .")
	if err := runCommandLogOutput(cmd); err != nil {
		return err
	}

	Log("finished building, starting up sandbox on port " + port)
	Log("run requests in another tab")
	cmd = exec.Command("bash", "-c", "docker run --rm -p "+port+":"+port+" projections")
	if err := runCommandLogOutput(cmd); err != nil {
		return err
	}
	return nil
}

func runCommandLogOutput(cmd *exec.Cmd) error {
	pipe, _ := cmd.StderrPipe()
	if err := cmd.Start(); err != nil {
		return err
	}
	reader := bufio.NewReader(pipe)
	line, err := reader.ReadString('\n')
	for err == nil {
		Log(line)
		line, err = reader.ReadString('\n')
	}
	cmd.Wait()
	return nil
}
