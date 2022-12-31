package utility

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

func GetNameFromDirectory(directory string) string {
	var name string
	if directory != "." {
		nameSlice := strings.Split(directory, "/")
		name = nameSlice[len(nameSlice)-1]
	} else {
		path, err := os.Getwd()
		if err != nil {
			return ""
		}
		nameSlice := strings.Split(path, "/")
		name = nameSlice[len(nameSlice)-1]
	}
	return name
}

func CreateDirectory(directory string) error {
	if directory != "." {
		// return an error if app exists
		if _, err := os.Stat(directory); !os.IsNotExist(err) {
			return errors.New(fmt.Sprintf("%s already exists", directory))
		}

		err := os.Mkdir(directory, 0750)
		if err != nil && !os.IsExist(err) {
			return err
		}
	}
	return nil
}

// given a directory and config interface, creates / reads .strapprc and sets config values
func SetupStrapprcConfig(directory string, config interface{}) error {
	if _, err := os.Stat(directory + "/.strapprc"); os.IsNotExist(err) {
		_, err = os.Create(directory + "/.strapprc")
		if err != nil {
			return err
		}
	}

	// get conf from conf file
	viper.SetConfigName(".strapprc")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(directory)
	if err := viper.MergeInConfig(); err != nil {
		return err
	}
	if err := viper.Unmarshal(config); err != nil {
		return err
	}
	return nil
}

func WriteConfigToStrapprc(directory string, config interface{}) error {
	jsonConfig, err := yaml.Marshal(config)
	if err != nil {
		return err
	}
	err = viper.ReadConfig(bytes.NewBuffer([]byte(jsonConfig)))
	if err != nil {
		return err
	}
	err = viper.WriteConfig()
	if err != nil {
		return err
	}
	return nil
}
