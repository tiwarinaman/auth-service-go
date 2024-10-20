package config

import (
	"errors"
	"github.com/spf13/viper"
	"log"
)

func Set() {

	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath("config")

	if err := viper.ReadInConfig(); err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if errors.As(err, &configFileNotFoundError) {
			log.Fatal("Error reading the configs")
		}
	}

	err := viper.Unmarshal(&configuration)
	if err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}

}
