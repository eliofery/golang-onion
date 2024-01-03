package main

import (
	"fmt"
	"github.com/eliofery/golang-angular/pkg/config"
	"github.com/eliofery/golang-angular/pkg/config/godotenv"
	"github.com/eliofery/golang-angular/pkg/config/viperr"
	"github.com/spf13/viper"
	"os"
)

func main() {
	config.Init(godotenv.New(".env"))
	fmt.Println(os.Getenv("SERVER_URL"))

	config.Init(viperr.New())
	fmt.Println(viper.GetString("server.url"))
}
