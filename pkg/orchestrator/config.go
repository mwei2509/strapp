package orchestrator

import (
	"bytes"
	"strings"

	"gopkg.in/yaml.v2"

	"github.com/spf13/viper"
)

type Config struct {
	Services   []Service  `mapstructure:"services"`
	Datastores []struct{} `mapstructure:"datastores" yaml:"datastores,omitempty"`
	Cicd       string     `mapstructure:"cicd"`
}

func (o *Orchestrator) setConfig() error {
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
	o.Config.Cicd = strings.ToLower(o.Config.Cicd)
	o.setServiceConfig()

	// rewriting config with format
	config, _ := yaml.Marshal(o.Config)
	viper.ReadConfig(bytes.NewBuffer([]byte(config)))
	viper.WriteConfig()

	return nil
}

func (o *Orchestrator) setServiceConfig() {
	// if this is not a single app
	// service directories will be nested in the main directory
	isSingleApp := len(o.Config.Services) == 1

	// port tracking
	for i, v := range o.Config.Services {
		directory := o.Directory
		if !isSingleApp {
			directory = o.Directory + "/" + v.Name
		}
		o.Config.Services[i].Directory = directory
		o.Config.Services[i].IsRootApp = isSingleApp
		o.Config.Services[i].Type = strings.ToLower(v.Type)
		o.Config.Services[i].Port = assignPort(o.Config.Services[i].Type, o.Config.Services[i].Port)
		o.Config.Services[i].Name = strings.ToLower(v.Name)
		o.Config.Services[i].Language = strings.ToLower(v.Language)
		o.Config.Services[i].Framework = strings.ToLower(v.Framework)
		o.Config.Services[i].Css = strings.ToLower(v.Css)
		o.Config.Services[i].StateManagement = strings.ToLower(v.StateManagement)
		o.Config.Services[i].Deployment = strings.ToLower(v.Deployment)
		o.Config.Services[i].StateManagement = strings.ToLower(v.StateManagement)
		o.Config.Services[i].Orm = strings.ToLower(v.Orm)

		for id, v := range v.Datastores {
			o.Config.Services[i].Datastores[id] = strings.ToLower(v)
		}
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
		if port != 0 {
			ports.apiPorts = append(ports.apiPorts, port)
		}
		// take last port + increment + push
		if len(ports.apiPorts) == 0 {
			port = 3000
		} else {
			port = ports.apiPorts[len(ports.apiPorts)-1] + 1
		}
		ports.apiPorts = append(ports.apiPorts, port)
	case "frontend":
		if port != 0 {
			ports.webPorts = append(ports.webPorts, port)
		}
		if len(ports.webPorts) == 0 {
			port = 8000
		} else {
			port = ports.webPorts[len(ports.webPorts)-1] + 1
		}
		ports.webPorts = append(ports.webPorts, port)
	case "database":
		if port != 0 {
			ports.dbPorts = append(ports.dbPorts, port)
		}
		if len(ports.dbPorts) == 0 {
			port = 5432
		} else {
			port = ports.dbPorts[len(ports.dbPorts)-1] + 1
		}
		ports.dbPorts = append(ports.dbPorts, port)
	}
	return port
}

// write after setting
func writeConf(config Config) error {
	return nil
}
