package sqlite

import (
	"database/sql"
	"github.com/eliofery/golang-onion/pkg/config"
	"github.com/eliofery/golang-onion/pkg/database"
	"github.com/gofiber/fiber/v3/log"
	_ "github.com/mattn/go-sqlite3"
)

// Storage база данных sqlite
// Пример: database.Connect(sqlite.New(config))
// config - godotenv, viperr
type Storage interface {
	database.Database
}

type storage struct {
	Path string
}

func New(config config.Config) Storage {
	log.Info("инициализация базы данных Sqlite")

	return &storage{
		Path: config.Get("SQLITE_PATH"),
	}
}

func (s *storage) Init() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", s.Path)
	if err != nil {
		return nil, database.ErrConnectDB
	}

	if err = db.Ping(); err != nil {
		return nil, database.ErrConnectDB
	}

	return db, nil
}
