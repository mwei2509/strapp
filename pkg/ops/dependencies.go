/*
Checks for necessary installations
*/
package ops

func GetInstallNeeded() []string {
	var installNeeded []string
	if !checkDockerExists() {
		installNeeded = append(installNeeded, "docker")
	}
	return installNeeded
}

func InstallDependencies() {

}
