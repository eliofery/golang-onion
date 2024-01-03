package viperr

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v3/log"
	"github.com/spf13/viper"
	"strings"
)

var (
	ErrNotFound = errors.New("конфигурационный файл не найден")
	ErrParse    = errors.New("не удалось обработать конфигурационный файл")
)

const (
	defaultConfigName = "config"
	defaultConfigType = "yml"
	defaultConfigPath = "internal/config"
)

// Viper загрузка конфигураций из yml файлов
// Пример: config.Init(viperr.New())
type Viper struct {
	configName  string
	configType  string
	configPaths []string
}

func New(configName ...string) *Viper {
	name := defaultConfigName
	if len(configName) > 0 {
		name = configName[0]
	}

	return &Viper{
		configName:  name,
		configType:  defaultConfigType,
		configPaths: []string{defaultConfigPath},
	}
}

func (v *Viper) AddConfigType(configType string) *Viper {
	v.configType = configType

	return v
}

func (v *Viper) AddConfigPath(configPath ...string) *Viper {
	v.configPaths = append(v.configPaths, configPath...)

	return v
}

func (v *Viper) Load() error {
	op := "viperr.Load"

	viper.SetConfigName(v.configName)
	viper.SetConfigType(v.configType)
	for _, configPath := range v.configPaths {
		viper.AddConfigPath(configPath)
	}

	if err := viper.ReadInConfig(); err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError

		log.Error(fmt.Errorf("%s: %w", op, err))
		if errors.As(err, &configFileNotFoundError) {
			return ErrNotFound
		}

		return ErrParse
	}

	return nil
}

func (v *Viper) Get(key string) string {
	key = formatter(key)

	return viper.GetString(key)
}

func (v *Viper) GetAny(key string) any {
	key = formatter(key)

	return viper.Get(key)
}

// Formatter изменяет строку под конфигурацию
// Пример: SERVER_PORT -> server.port
func formatter(key string) string {
	key = strings.ToLower(key)
	key = strings.ReplaceAll(key, "_", ".")

	return key
}
