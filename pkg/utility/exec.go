package utility

import (
	"bufio"
	"fmt"
	"os/exec"
)

type Cmd struct {
	Dir string
	Cmd string
}

func CheckCommandExists(command string) bool {
	out, err := exec.Command("command", "-v", command).Output()
	fmt.Print(string(out))
	if err != nil {
		return false
	}
	return true
}

// params argument
// [0] = directory
func RunCommand(command string, params ...string) (string, error) {
	directory := ""
	if len(params) > 0 {
		directory = params[0]
	}
	cmd := exec.Command("bash", "-c", command)
	if directory != "" {
		cmd.Dir = directory
	}
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(out), nil
}

func RunCommandLogOutput(command string, params ...string) error {
	directory := ""
	if len(params) > 0 {
		directory = params[0]
	}
	cmd := exec.Command("bash", "-c", command)
	if directory != "" {
		cmd.Dir = directory
	}
	r, _ := cmd.StdoutPipe()
	//rUse the same pipe for standard error
	cmd.Stderr = cmd.Stdout

	// Make a new channel which will be used to ensure we get all output
	done := make(chan struct{})

	// Create a scanner which scans r in a line-by-line fashion
	scanner := bufio.NewScanner(r)
	// Use the scanner to scan the output line by line and log it
	// It's running in a goroutine so that it doesn't block
	go func() {

		// Read line by line and process it
		for scanner.Scan() {
			line := scanner.Text()
			log.Log(line)
		}

		// We're all done, unblock the channel
		done <- struct{}{}

	}()
	// Start the command and check for errors
	if err := cmd.Start(); err != nil {
		return err
	}

	// Wait for all output to be processed
	<-done

	// Wait for the command to finish
	if err := cmd.Wait(); err != nil {
		return err
	}
	return nil
}
