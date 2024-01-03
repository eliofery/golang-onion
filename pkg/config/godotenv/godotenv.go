package godotenv

import (
	"errors"
	"fmt"
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
type GoDotEnv struct {
	configName []string
}

func New(configName ...string) *GoDotEnv {
	if len(configName) == 0 {
		configName = append(configName, defaultConfigName)
	}

	return &GoDotEnv{
		configName: configName,
	}
}

func (g *GoDotEnv) Load() error {
	op := "godotenv.Load"

	if err := godotenv.Load(g.configName...); err != nil {
		log.Error(fmt.Errorf("%s: %w", op, err))
		return ErrNotFound
	}

	return nil
}

func (g *GoDotEnv) Get(key string) string {
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
