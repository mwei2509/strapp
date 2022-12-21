package orchestrator

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	"gopkg.in/yaml.v2"

	"github.com/mwei2509/strapp/pkg/ops"
	"github.com/spf13/viper"
)

type Config struct {
	Services      []Service  `mapstructure:"services"`
	Databases     []Database `mapstructure:"databases" yaml:"databases,omitempty"`
	Cicd          struct{}   `mapstructure:"cicd"`
	DockerCompose ops.DockerCompose
}

func (o *Orchestrator) setConfig() error {
	o.Log(o.Directory)
	if _, err := os.Stat(o.Directory + "/.strapprc"); os.IsNotExist(err) {
		o.Log("creating config")
		_, err = os.Create(o.Directory + "/.strapprc")
		if err != nil {
			return err
		}
	}

	// get conf from conf file
	viper.SetConfigName(".strapprc")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(o.Directory)
	err := viper.MergeInConfig()

	if err != nil {
		return err
	}

	// write conf to struct
	if err = viper.Unmarshal(&o.Config); err != nil {
		return err
	}

	// formatting
	// o.Config.Cicd = strings.ToLower(o.Config.Cicd)
	o.setServicesConfig()
	o.setDatabasesConfig()
	o.setDockerComposeConfig()

	// rewriting config with format
	config, _ := yaml.Marshal(o.Config)
	viper.ReadConfig(bytes.NewBuffer([]byte(config)))
	viper.WriteConfig()

	return nil
}

func (o *Orchestrator) setServicesConfig() {
	// if this is not a single app
	// service directories will be nested in the main directory
	isSingleApp := len(o.Config.Services) == 1

	// port tracking
	for i, v := range o.Config.Services {
		directory := o.Directory
		if !isSingleApp {
			directory = o.Directory + "/" + v.Name
		}
		serviceType := strings.ToLower(v.Type)
		o.Config.Services[i].Directory = directory
		o.Config.Services[i].IsRootApp = isSingleApp
		o.Config.Services[i].Type = serviceType
		port := assignPort(serviceType, v.Port)
		o.Config.Services[i].Port = port
		if strings.ToLower(serviceType) == "api" {
			o.Config.Services[i].DebuggerPort = port + 2000
		}
		o.Config.Services[i].Name = strings.ToLower(v.Name)
		o.Config.Services[i].Language = strings.ToLower(v.Language)
		o.Config.Services[i].Framework = strings.ToLower(v.Framework)
		o.Config.Services[i].Css = strings.ToLower(v.Css)
		o.Config.Services[i].StateManagement = strings.ToLower(v.StateManagement)
		o.Config.Services[i].Deployment = strings.ToLower(v.Deployment)
		o.Config.Services[i].StateManagement = strings.ToLower(v.StateManagement)
		o.Config.Services[i].Orm = strings.ToLower(v.Orm)

		for id, v := range v.Databases {
			o.Config.Services[i].Databases[id] = strings.ToLower(v)
		}
	}
}

func (o *Orchestrator) setDatabasesConfig() {
	for i, v := range o.Config.Databases {
		o.Config.Databases[i].Type = strings.ToLower(v.Type)
		o.Config.Databases[i].Port = assignPort("database", o.Config.Databases[i].Port)
		o.Config.Databases[i].Name = strings.ToLower(v.Name)
	}
}

func (o *Orchestrator) setDockerComposeConfig() {
	o.Config.DockerCompose.Services = make(map[string]ops.DockerComposeService)
	o.Config.DockerCompose.Volumes = make(map[string]ops.DockerComposeVolume)
	for _, v := range o.Config.Databases {
		dbService := ops.DockerComposeService{}
		dbName := strings.ToLower(v.Name)
		if v.Type == "postgresql" {
			dbService.Image = "postgres:14.1-alpine"
			dbService.Restart = "always"
			dbService.Environment = make(map[string]string)
			dbService.Environment["POSTGRES_USER"] = "postgres"
			dbService.Environment["POSTGRES_PASSWORD"] = "password"
			dbService.Environment["POSTGRES_DB"] = o.Name + "_development"
			dbService.Ports = []string{fmt.Sprint(v.Port) + ":" + fmt.Sprint(v.Port)}
			// add a volume
			o.Config.DockerCompose.Volumes[dbName] = ops.DockerComposeVolume{
				Driver: "local",
			}
			dbService.Volumes = []string{
				dbName + ":/var/lib/postgresql/data",
				"./docker/" + dbName + "/init.sql:docker-entrypoint-initdb.d/create_tables.sql",
			}
		}
		o.Config.DockerCompose.Services[dbName] = dbService
	}
	for _, v := range o.Config.Services {
		service := ops.DockerComposeService{}
		name := strings.ToLower(v.Name)
		service.ContainerName = name
		service.Build.Context = "."
		service.Build.Dockerfile = v.Directory + "/Dockerfile"
		service.Environment = make(map[string]string)
		service.Environment["PORT"] = fmt.Sprint(v.Port)
		service.Ports = []string{fmt.Sprint(v.Port) + ":" + fmt.Sprint(v.Port)}
		if v.DebuggerPort != 0 {
			service.Ports = append(service.Ports, fmt.Sprint(v.DebuggerPort)+":"+fmt.Sprint(v.DebuggerPort))
		}
		service.Volumes = []string{
			v.Directory + ":/app:delegated",
		}
		o.Config.DockerCompose.Services[name] = service
	}
}

// ports
// api ports in 3000's
// api debugger ports in 4000's
// web ports 8000's
// db ports 5000's
type Ports struct {
	apiPorts []int64
	webPorts []int64
	dbPorts  []int64
}

var ports Ports

func assignPort(portType string, requested int64) int64 {
	var port int64
	if requested != 0 {
		port = requested
	}
	switch portType {
	case "api":
		if port == 0 {
			if len(ports.apiPorts) == 0 {
				port = 3000
			} else {
				port = ports.apiPorts[len(ports.apiPorts)-1] + 1
			}
		}
		ports.apiPorts = append(ports.apiPorts, port)
	case "frontend":
		if port == 0 {
			if len(ports.webPorts) == 0 {
				port = 8000
			} else {
				port = ports.webPorts[len(ports.webPorts)-1] + 1
			}
		}
		ports.webPorts = append(ports.webPorts, port)
	case "database":
		if port == 0 {
			if len(ports.dbPorts) == 0 {
				port = 5432
			} else {
				port = ports.dbPorts[len(ports.dbPorts)-1] + 1
			}
		}
		ports.dbPorts = append(ports.dbPorts, port)
	}
	return port
}

// write after setting
func writeConf(config Config) error {
	return nil
}
