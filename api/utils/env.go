package utils

import (
	"ffcs/api/constants"
	"log"
	"strings"

	"github.com/spf13/viper"
)

func ImportEnvs() {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	viper.AddConfigPath(".")
	viper.SetDefault("PORT", 3000)
	viper.SetDefault("MIGRATE", false)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil { // reading the .env
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Println("[ENV] .env file not found")
		} else {
			log.Fatalf("[ENV] Error Reading the program")
		}

	}

	for _, element := range constants.ENV {
		if viper.GetString(element) == "" {
			log.Fatalf("[ENV] Envionment variable is not present %s", element)
		}
	}
	constants.Branches = strings.Split(viper.GetString("BRANCHES"), ",")
}
