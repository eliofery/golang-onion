package sqlite

import (
	"database/sql"
	"github.com/eliofery/golang-angular/pkg/config"
	"github.com/eliofery/golang-angular/pkg/database"
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
	db, err := sql.Open("sqlite3", s.Path)
	if err != nil {
		return nil, database.ErrConnectDB
	}

	if err = db.Ping(); err != nil {
		return nil, database.ErrConnectDB
	}

	return db, nil
}
