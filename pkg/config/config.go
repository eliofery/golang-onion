package config

import (
	"github.com/gofiber/fiber/v3/log"
)

type Config interface {
	// Load загрузка конфигурации
	Load() error

	// Get получение конфигурации
	Get(key string) string
}

// Init инициализация конфигурации
// Пример: config.Init(viperr.New())
func Init(config Config) (Config, error) {
	if err := config.Load(); err != nil {
		return nil, err
	}

	return config, nil
}

// MustInit инициализация конфигурации с обработкой ошибок
func MustInit(config Config) Config {
	log.Info("инициализация конфигурации")

	conf, err := Init(config)
	if err != nil {
		log.Fatal(err)
	}

	return conf
}
