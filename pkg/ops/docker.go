package ops

import (
	"fmt"
	"os/exec"
)

func checkDockerExists() bool {
	out, err := exec.Command("command", "-v", "docker").Output()
	Log(fmt.Sprintf("%s\n", out))
	if err != nil {
		return false
	}
	return true
}

func installDocker() error {
	return nil
}
