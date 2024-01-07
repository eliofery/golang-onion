package database

import (
	"database/sql"
	"errors"
	"github.com/gofiber/fiber/v3/log"
)

var (
	ErrConnectDB = errors.New("не удалось подключиться к базе данных")
)

type Database interface {
	// Init инициализация БД
	Init() (*sql.DB, error)
}

// Connect подключение к БД
// Пример: database.Connect(postgres.New(config))
// config - godotenv, viperr
func Connect(driver Database) (*sql.DB, error) {
	db, err := driver.Init()
	if err != nil {
		return nil, err
	}

	return db, nil
}

// MustConnect подключение к БД с обработкой ошибок
func MustConnect(driver Database) *sql.DB {
	log.Info("подключение к БД")

	db, err := Connect(driver)
	if err != nil {
		log.Fatalf("%s: %s", ErrConnectDB, err)
	}

	return db
}
