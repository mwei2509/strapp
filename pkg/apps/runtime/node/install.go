package node

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

var ErrInstallation = errors.New("installation error")

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

func (n Node) isNodeVersion() bool {
	out, err := exec.Command("node", "--version").Output()
	if err != nil {
		// install npm
		return false
	}
	if !strings.Contains(fmt.Sprintf("%s\n", out), "v"+settings.Version) {
		return false
	}
	return true
}

func (n Node) installNodeVersion() error {
	file, err := os.CreateTemp("", "nvmshell-*.sh")
	if err != nil {
		n.Fatal(err)
		return ErrInstallation
	}

	defer os.Remove(file.Name())

	_, err = file.WriteString(nvmShell)
	if err != nil {
		n.Fatal(err)
		return ErrInstallation
	}

	out, err := exec.Command("/bin/sh", file.Name(), settings.Version).Output()
	if err != nil {
		n.Fatal(err)
		return ErrInstallation
	}

	n.Log(fmt.Sprintf("%s\n", out))
	return nil
}

func (n Node) setNodeDependencies() error {
	if !n.isNodeVersion() {
		// install node
		if err := n.installNodeVersion(); err != nil {
			return err
		}
	}
	return nil
}

func (n Node) installNvm() error {
	cmd := "curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.39.2/install.sh | bash"
	out, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		n.Fatal(err)
		return ErrInstallation
	}

	n.Log(fmt.Sprintf("%s\n", out))
	return nil
}
