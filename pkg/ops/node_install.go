package ops

import (
	"errors"
	"os"
	"strings"

	u "github.com/mwei2509/strapp/pkg/utility"
)

var ErrInstallation = errors.New("installation error")

type NodeSettings struct {
	Version string
}

var settings NodeSettings = NodeSettings{
	Version: "18",
}

var nvmShell string = `
if [ "$#" -ne 1 ]; then
	echo "incorrect usage"
	echo "Usage:"
	echo "  SetNodeVersion <VERSION>"
	exit -1
fi

NODE_VERSION=$1

export NVM_DIR="${HOME}/.nvm"
[ -s "${NVM_DIR}/nvm.sh" ] && . "$NVM_DIR/nvm.sh"  # This loads nvm

nvm use ${NODE_VERSION} || nvm install ${NODE_VERSION}
`

func isNodeVersion() bool {
	out, err := u.RunCommand("node --version")
	if err != nil {
		return false
	}
	if !strings.Contains(out, "v"+settings.Version) {
		return false
	}
	return true
}

func installNodeVersion() error {
	file, err := os.CreateTemp("", "nvmshell-*.sh")
	if err != nil {
		log.Fatal(err)
		return ErrInstallation
	}

	defer os.Remove(file.Name())

	_, err = file.WriteString(nvmShell)
	if err != nil {
		log.Fatal(err)
		return ErrInstallation
	}

	out, err := u.RunCommand("/bin/sh " + file.Name() + " " + settings.Version)
	if err != nil {
		log.Fatal(err)
		return ErrInstallation
	}
	log.Log(out)
	return nil
}

func InstallNode() error {
	if !isNodeVersion() {
		// install node
		if err := installNodeVersion(); err != nil {
			return err
		}
	}
	return nil
}

func NpmInstall(directory string) error {
	if err := u.RunCommandLogOutput("npm install", directory); err != nil {
		return err
	}
	return nil
}

func NpmInstallPkg(directory string, pkg string) error {
	out, err := u.RunCommand("npm install "+pkg, directory)
	if err != nil {
		return err
	}
	log.Log(out)
	return nil
}

func installNvm() error {
	cmd := "curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.39.2/install.sh | bash"

	if err := u.RunCommandLogOutput(cmd); err != nil {
		log.Fatal(err)
		return ErrInstallation
	}
	return nil
}
