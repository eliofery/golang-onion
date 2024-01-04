package main

import (
	"github.com/eliofery/golang-angular/pkg/config"
	"github.com/eliofery/golang-angular/pkg/config/godotenv"
	"github.com/eliofery/golang-angular/pkg/config/viperr"
	"github.com/eliofery/golang-angular/pkg/database"
	"github.com/eliofery/golang-angular/pkg/database/postgres"
	"github.com/eliofery/golang-angular/pkg/database/sqlite"
	"github.com/gofiber/fiber/v3/log"
	"github.com/spf13/viper"
	"os"
)

func main() {
	// Тест godotenv
	env, err := config.Init(godotenv.New(".env"))
	if err == nil {
		log.Info(os.Getenv("SERVER_URL"))
	}

	// Тест viperr
	yml, err := config.Init(viperr.New())
	if err == nil {
		log.Info(viper.GetString("server.url"))
	}

	// Тест sqlite
	_, err = database.Connect(sqlite.New(env))
	if err == nil {
		log.Info("подключение БД sqlite")
	}

	// Тест postgres
	_, err = database.Connect(postgres.New(yml))
	if err == nil {
		log.Info("подключение БД postgres")
	}
}
