package config

import (
	"log"

	"github.com/spf13/viper"
)

var EnvConfigs *envConfigs

type envConfigs struct {
	DiscordToken string `mapstructure:"DISCORD_TOKEN"`
}

// Call to load the variables from env
func LoadConfig() *envConfigs {
	viper.AddConfigPath("../")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error reading env file", err)
	}

	if err := viper.Unmarshal(&EnvConfigs); err != nil {
		log.Fatal(err)
	}

	return EnvConfigs
}
