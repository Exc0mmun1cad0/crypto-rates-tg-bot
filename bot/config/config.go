package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	BatchSize int    `json:"batchSize"`
	Host      string `json:"host"`
}

func NewBotConfig(configPath string) (*Config, error) {
	var cfg Config

	err := cleanenv.ReadConfig(configPath, &cfg)
	if err != nil {
		return nil, fmt.Errorf("cannot read config")
	}

	return &cfg, nil
}
