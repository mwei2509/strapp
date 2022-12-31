/*
Checks for necessary installations
*/
package ops

import (
	u "github.com/mwei2509/strapp/pkg/utility"
)

func GetInstallNeeded() []string {
	var installNeeded []string
	if !u.CheckCommandExists("docker") {
		installNeeded = append(installNeeded, "docker")
	}
	return installNeeded
}

func InstallDependencies() {

}
