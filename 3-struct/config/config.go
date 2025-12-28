package config

import (
	"os"
)

type AppConfig struct {
	key string
}

func NewConfig() *AppConfig {
	var key = os.Getenv("KEY")

	if key == "" {
		panic("Переменная окружения KEY отсутствует")
	}

	return &AppConfig{
		key: key,
	}
}

func (cfg *AppConfig) GetKey() string {
	return cfg.key
}
