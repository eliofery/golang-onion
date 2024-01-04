package postgres

import (
	"database/sql"
	"fmt"
	"github.com/eliofery/golang-angular/pkg/config"
	"github.com/eliofery/golang-angular/pkg/database"
	"github.com/gofiber/fiber/v3/log"
	_ "github.com/jackc/pgx/v5/stdlib"
	"strconv"
)

const (
	portDefault = 5432
)

// Storage база данных postgres
// Пример: database.Connect(postgres.New(config))
// config - godotenv, viperr
type Storage struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
	SSLMode  string
}

func New(config config.Config) *Storage {
	port, err := strconv.Atoi(config.Get("POSTGRES_PORT"))
	if err != nil {
		port = portDefault
	}

	return &Storage{
		Host:     config.Get("POSTGRES_HOST"),
		Port:     port,
		User:     config.Get("POSTGRES_USER"),
		Password: config.Get("POSTGRES_PASSWORD"),
		Database: config.Get("POSTGRES_DATABASE"),
		SSLMode:  config.Get("POSTGRES_SSLMODE"),
	}
}

func (s *Storage) Init() (*sql.DB, error) {
	op := "postgres.Connect"

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", s.Host, s.Port, s.User, s.Password, s.Database, s.SSLMode)

	db, err := sql.Open("pgx", dsn)
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
