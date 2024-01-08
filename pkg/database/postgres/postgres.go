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
type Storage interface {
	database.Database
}

type storage struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
	SSLMode  string
}

func New(config config.Config) Storage {
	log.Info("инициализация базы данных Postgres")

	port, err := strconv.Atoi(config.Get("POSTGRES_PORT"))
	if err != nil {
		port = portDefault
	}

	return &storage{
		Host:     config.Get("POSTGRES_HOST"),
		Port:     port,
		User:     config.Get("POSTGRES_USER"),
		Password: config.Get("POSTGRES_PASSWORD"),
		Database: config.Get("POSTGRES_DATABASE"),
		SSLMode:  config.Get("POSTGRES_SSLMODE"),
	}
}

func (s *storage) Init() (*sql.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", s.Host, s.Port, s.User, s.Password, s.Database, s.SSLMode)

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
