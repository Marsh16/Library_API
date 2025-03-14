package main

// go mod init vp_week11_echo
// GO111MODULE=on go get github.com/labstack/echo/v4
// go get github.com/tkanos/gonfig
// go get -u github.com/go-sql-driver/mysql
//go get golang.org/x/crypto/bcrypt
//go get github.com/dgrijalva/jwt-go
//go get github.com/labstack/echo/v4/middleware
//go get github.com/go-playground/validator
//go get github.com/joho/godotenv

import (
	"library_api/db"
	"library_api/routes"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	db.Init()
	e := routes.Init()
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	  }
	baseURL := os.Getenv("BASE_URL")
	if baseURL == "" {
		e.Logger.Fatal("BASE_URL config is required")
	}
	e.Logger.Fatal(e.Start(baseURL))
}
