/*
Checks for necessary installations
*/
package app

func getInstallNeeded() []string {
	var installNeeded []string
	if !checkDockerExists() {
		installNeeded = append(installNeeded, "docker")
	}
	return installNeeded
}

func installDependencies() {

}
