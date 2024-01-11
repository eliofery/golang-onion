package viperr

import (
	"errors"
	"github.com/eliofery/golang-onion/pkg/config"
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

// Viperr загрузка конфигураций из yml файлов
// Пример: config.Init(viperr.New())
type Viperr interface {
	config.Config
	AddConfigType(configType string) *viperr
	AddConfigPath(configPath ...string) *viperr
	GetAny(key string) any
}

type viperr struct {
	configName  string
	configType  string
	configPaths []string
}

func New(configName ...string) Viperr {
	log.Info("инициализация конфигурации viperr")

	name := defaultConfigName
	if len(configName) > 0 {
		name = configName[0]
	}

	return &viperr{
		configName:  name,
		configType:  defaultConfigType,
		configPaths: []string{defaultConfigPath},
	}
}

func (v *viperr) AddConfigType(configType string) *viperr {
	v.configType = configType

	return v
}

func (v *viperr) AddConfigPath(configPath ...string) *viperr {
	v.configPaths = append(v.configPaths, configPath...)

	return v
}

func (v *viperr) Init() error {
	viper.SetConfigName(v.configName)
	viper.SetConfigType(v.configType)
	for _, configPath := range v.configPaths {
		viper.AddConfigPath(configPath)
	}

	if err := viper.ReadInConfig(); err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError

		if errors.As(err, &configFileNotFoundError) {
			return ErrNotFound
		}

		return ErrParse
	}

	return nil
}

func (v *viperr) Get(key string) string {
	key = formatter(key)

	return viper.GetString(key)
}

func (v *viperr) GetAny(key string) any {
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
