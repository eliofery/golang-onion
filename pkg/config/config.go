package config

import (
	"fmt"
	"github.com/gofiber/fiber/v3/log"
)

type Config interface {
	// Load загрузка конфигурации
	Load() error

	// Get получение конфигурации
	Get(key string) string
}

// Init инициализация конфигурации
func Init(config Config) (Config, error) {
	op := "config.Init"

	if err := config.Load(); err != nil {
		log.Error(fmt.Errorf("%s: %w", op, err))
		return nil, err
	}

	return config, nil
}

// MustInit инициализация конфигурации с обработкой ошибок
func MustInit(config Config) Config {
	conf, err := Init(config)
	if err != nil {
		log.Fatal(err)
	}

	return conf
}
