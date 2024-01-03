package main

import (
	"fmt"
	"github.com/eliofery/golang-angular/pkg/config"
	"github.com/eliofery/golang-angular/pkg/config/godotenv"
	"os"
)

func main() {
	config.Init(godotenv.New(".env"))
	fmt.Print(os.Getenv("SERVER_URL"))
}
