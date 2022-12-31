package ops

import (
	"fmt"
	"os/exec"
)

func LoginGithubApi() bool {
	out, err := exec.Command("gh", "auth", "status").Output()
	fmt.Printf("%s\n", out)
	if err != nil {
		return false
	}
	return true
}
