package config

import (
	"github.com/tkanos/gonfig"
	"fmt"
	"os"
	"main/models"
)

func GetConfig(params ...string) models.Configuration {
	configuration := models.Configuration{}

	// Build & retrieve the configuration file based on environment variable
	env := os.Getenv("ENV")
	fileName := fmt.Sprintf("./config/configuration.%s.json", env)
	gonfig.GetConf(fileName, &configuration)
	return configuration
}