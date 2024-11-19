package config

import (
	"github.com/caarlos0/env/v10"
)

type Config struct {
	ConnectionString string `env:"POSTGRES_CONNECTION_STRING,required"`
	ServerAddr       string `env:"SERVER_ADDR" envDefault:"0.0.0.0:8080"`
}

func New() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
