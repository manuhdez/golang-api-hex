package config

import (
	"github.com/kelseyhightower/envconfig"
	"log"
)

type Config struct {
	App ApiConfig
	Db  DbConfig
}

func GetEnv() Config {
	var config Config
	err := envconfig.Process("", &config)
	if err != nil {
		log.Fatal(err.Error())
	}
	return config
}
