package config

import (
	"log"

	"github.com/spf13/viper"
)

var Configuration Config

func Set() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("src/config")
	viper.AddConfigPath("$HOME/.appname")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("error reading the config")
	}

	err := viper.Unmarshal(&Configuration)
	if err != nil {
		log.Fatal("unable to decode into struct", err)
	}
}
