package configs

import (
	"github.com/spf13/viper"
)

var config *Configuration

type Configuration struct {
	Database Database
}

func Get() *Configuration {
	return config
}

func Setup() (err error) {
	var c Configuration

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	if err := viper.Unmarshal(&c); err != nil {
		return err
	}

	config = &c
	return
}
