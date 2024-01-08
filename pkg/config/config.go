package config

import (
	"github.com/gofiber/fiber/v3/log"
)

type Config interface {
	// Init загрузка конфигурации
	Init() error

	// Get получение конфигурации
	Get(key string) string
}

// Load загрузка конфигурации
// Пример: config.Load(viperr.New())
func Load(config Config) (Config, error) {
	log.Info("загрузка конфигурации")

	if err := config.Init(); err != nil {
		return nil, err
	}

	return config, nil
}

// MustInit инициализация конфигурации с обработкой ошибок
func MustInit(config Config) Config {
	conf, err := Load(config)
	if err != nil {
		log.Fatal(err)
	}

	return conf
}
