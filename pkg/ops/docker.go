package ops

import (
	"fmt"
	"io/ioutil"
	"os/exec"

	"gopkg.in/yaml.v2"
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

var dockerTemplate = `

`

type DockerCompose struct {
	Version  string                          `yaml:"version"`
	Volumes  map[string]DockerComposeVolume  `yaml:"volumes,omitempty"`
	Services map[string]DockerComposeService `yaml:"services,omitempty"`
}

type DockerComposeVolume struct {
	Driver string `yaml:"driver,omitempty"`
}

type DockerComposeService struct {
	ContainerName string            `yaml:"container_name,omitempty"`
	Image         string            `yaml:"image,omitempty"`
	Restart       string            `yaml:"restart,omitempty"`
	Environment   map[string]string `yaml:"environment,omitempty"`
	Ports         []string          `yaml:"ports,omitempty"`
	Volumes       []string          `yaml:"volumes,omitempty"`
	Build         struct {
		Context    string `yaml:"context,omitempty"`
		Dockerfile string `yaml:"dockerfile,omitempty"`
	} `yaml:"yaml,omitempty"`
	Command string `yaml:"command,omitempty"`
}

func (dc *DockerCompose) ReadDockerCompose() error {
	dockerComposeFile, err := ioutil.ReadFile("docker-compose.yml")
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(dockerComposeFile, &dc)
	if err != nil {
		return err
	}
	return nil
}

func (dc *DockerCompose) WriteDockerCompose(directory string) error {
	dockerCompose, err := yaml.Marshal(dc)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(directory+"/docker-compose.yml", dockerCompose, 0644)
	if err != nil {
		return err
	}
	return nil
}
