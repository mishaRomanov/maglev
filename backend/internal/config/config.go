package config

import (
	"github.com/caarlos0/env/v9"
	log "github.com/sirupsen/logrus"
)

type GlobalConfig struct {
	RedisConfig
}

type RedisConfig struct {
	Port     string `env:"REDIS_PORT" envDefault:"6379"`
	Password string `env:"REDIS_PASSWORD" envDefault:""`
	DB       int    `env:"REDIS_DB" envDefault:"0"`
}

func NewFromEnv() *GlobalConfig {
	var envConfig GlobalConfig

	if parseErr := env.Parse(&envConfig); parseErr != nil {
		log.Fatal(parseErr)
	}
	return &envConfig
}
