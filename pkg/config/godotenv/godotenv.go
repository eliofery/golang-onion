package godotenv

import (
	"errors"
	"github.com/eliofery/golang-angular/pkg/config"
	"github.com/gofiber/fiber/v3/log"
	"github.com/joho/godotenv"
	"os"
	"strings"
)

var (
	ErrNotFound = errors.New("не удалось загрузить переменные окружения из файла .env")
)

const (
	defaultConfigName = ".env"
)

// GoDotEnv загрузка конфигураций из переменного окружения
// Пример: config.Init(godotenv.New(".env"))
type GoDotEnv interface {
	config.Config
}

type goDotEnv struct {
	configName []string
}

func New(configName ...string) GoDotEnv {
	log.Info("инициализация конфигурации goDotEnv")

	if len(configName) == 0 {
		configName = append(configName, defaultConfigName)
	}

	return &goDotEnv{
		configName: configName,
	}
}

func (g *goDotEnv) Init() error {
	if err := godotenv.Load(g.configName...); err != nil {
		return ErrNotFound
	}

	return nil
}

func (g *goDotEnv) Get(key string) string {
	key = formatter(key)

	return os.Getenv(key)
}

// Formatter изменяет строку под конфигурацию
// Пример: server.port -> SERVER_PORT
func formatter(key string) string {
	key = strings.ToUpper(key)
	key = strings.ReplaceAll(key, ".", "_")

	return key
}
