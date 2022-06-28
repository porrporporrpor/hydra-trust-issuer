package config

import (
	"fmt"
	"log"
	"strings"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

var config Config

func Init() *Config {
	err := godotenv.Load()
	if err != nil {
		envFileNotFound := strings.Contains(err.Error(), "no such file or directory")
		if !envFileNotFound {
			log.Println(fmt.Sprintf("read config error: %v", err.Error()))
		} else {
			log.Println("use environment from OS")
		}
	}
	err = envconfig.Process("", &config)
	if err != nil {
		log.Println(fmt.Sprintf("parse config error: %v", err))
	}

	return &config
}

func GetConfig() *Config {
	return &config
}
