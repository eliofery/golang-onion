package database

import (
	"database/sql"
	"errors"
	"github.com/eliofery/golang-angular/internal/database"
	"github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v3/log"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/mattn/go-sqlite3"
	"github.com/pressly/goose/v3"
)

var (
	ErrSetDialect     = errors.New("не удалось установить используемую БД")
	ErrUpMigration    = errors.New("не удалось выполнить миграцию базы данной")
	ErrCurrentDialect = errors.New("не удалось определить используемую БД")
)

const (
	dirMigration = "migration"
)

// Migrate миграция базы данных
// Пример: database.Migrate(database.Connect(postgres.New(config.Init(viperr.New()))))
func Migrate(db *sql.DB) error {
	goose.SetBaseFS(database.EmbedMigration)
	defer goose.SetBaseFS(nil)

	if err := goose.SetDialect(getCurrentDialect(db)); err != nil {
		return ErrSetDialect
	}

	if err := goose.Up(db, dirMigration); err != nil {
		return ErrUpMigration
	}

	return nil
}

// GetCurrentDialect получение названия используемой базы данной
func getCurrentDialect(db *sql.DB) string {
	op := "database.getCurrentDialect"

	var dialect goose.Dialect

	switch db.Driver().(type) {
	case *mysql.MySQLDriver:
		dialect = goose.DialectMySQL
	case *sqlite3.SQLiteDriver:
		dialect = goose.DialectSQLite3
	case *stdlib.Driver:
		dialect = goose.DialectPostgres
	default:
		log.Errorf("%s: %s", op, ErrCurrentDialect)
	}

	return string(dialect)
}
