package app

import (
	"github.com/spf13/viper"
)

type AppConfig struct {
	Name      string `mapstructure:"name"`
	Type      string `mapstructure:"type"`
	Language  string `mapstructure:"language"`
	Framework string `mapstructure:"framework"`
}

type Config struct {
	App        AppConfig   `mapstructure:"app"`
	Apps       []AppConfig `mapstructure:"apps"`
	Deployment string      `mapstructure:"deployment"`
}

func getConf(directory string) (Config, error) {
	//
	// get conf from conf file
	viper.SetConfigName(".strapprc")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(directory)
	err := viper.MergeInConfig()
	if err != nil {
		return Config{}, err
	}

	// write conf to struct
	var config Config
	if err = viper.Unmarshal(&config); err != nil {
		return Config{}, err
	}
	return config, nil
}

// write after setting
func writeConf() error {
	return nil
}
