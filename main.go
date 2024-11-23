package main

// go mod init vp_week11_echo
// GO111MODULE=on go get github.com/labstack/echo/v4
// go get github.com/tkanos/gonfig
// go get -u github.com/go-sql-driver/mysql
//go get golang.org/x/crypto/bcrypt
//go get github.com/dgrijalva/jwt-go
//go get github.com/labstack/echo/v4/middleware
//go get github.com/go-playground/validator
import (
	"library_api/db"
	"library_api/helpers"
	"library_api/routes"
)

func main() {
	db.Init()
	e := routes.Init()
	baseURL := helpers.GetBaseURL()
	e.Logger.Fatal(e.Start(baseURL))
}
