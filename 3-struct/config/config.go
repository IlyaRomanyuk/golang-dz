package config

import (
	"os"
)

type AppConfig struct {
	key string
	url string
}

func NewConfig() *AppConfig {
	var key = os.Getenv("KEY")
	var url = os.Getenv("URL")

	if key == "" {
		panic("Переменная окружения KEY отсутствует")
	}

	if url == "" {
		panic("Переменная окружения URL отсутствует")
	}

	return &AppConfig{
		key: key,
		url: url,
	}
}

func (cfg *AppConfig) GetKey() string {
	return cfg.key
}

func (cfg *AppConfig) GetUrl() string {
	return cfg.url
}
