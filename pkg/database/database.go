package database

import (
	"database/sql"
	"errors"
	"fmt"
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
	op := "database.Connect"

	db, err := driver.Init()
	if err != nil {
		log.Error(fmt.Errorf("%s: %w", op, err))

		return nil, ErrConnectDB
	}

	return db, nil
}

// MustConnect подключение к БД с обработкой ошибок
func MustConnect(driver Database) *sql.DB {
	db, err := Connect(driver)
	if err != nil {
		log.Fatal(err)
	}

	return db
}
