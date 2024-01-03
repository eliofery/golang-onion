package sqlite

import (
	"database/sql"
	"fmt"
	"github.com/eliofery/golang-angular/pkg/config"
	"github.com/eliofery/golang-angular/pkg/database"
	"github.com/gofiber/fiber/v3/log"
	_ "github.com/mattn/go-sqlite3"
)

// Storage база данных sqlite
// Пример: database.Connect(sqlite.New(config))
// config - godotenv, viperr
type Storage struct {
	Path string
}

func New(config config.Config) *Storage {
	return &Storage{
		Path: config.Get("SQLITE_PATH"),
	}
}

func (s *Storage) Init() (*sql.DB, error) {
	op := "sqlite.Init"

	db, err := sql.Open("sqlite3", s.Path)
	if err != nil {
		log.Error(fmt.Errorf("%s: %w", op, err))
		return nil, database.ErrConnectDB
	}

	if err = db.Ping(); err != nil {
		log.Error(fmt.Errorf("%s: %w", op, err))
		return nil, database.ErrConnectDB
	}

	return db, nil
}
