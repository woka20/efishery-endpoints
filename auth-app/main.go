package main

import (
	"efishery-endpoints/auth-app/config"
	"efishery-endpoints/auth-app/handler/tokenization"
	"efishery-endpoints/auth-app/handler/user"
	"fmt"

	"github.com/kataras/iris/v12"
)

func main() {

	config.ConfigInit()

	apps := iris.New()

	userHandle := user.NewUserHandler()

	apps.Handle("GET", "/hello", userHandle.Hello)
	apps.Handle("POST", "/adduser", userHandle.AddNewUser)

	tokenHandle := tokenization.NewTokenHandler()
	apps.Handle("GET", "/token", tokenHandle.GetToken)
	apps.Handle("GET", "/claims", tokenHandle.GetClaims)

	listenPort := config.PORT
	fmt.Println("Server Ready!")

	apps.Listen(":" + listenPort)
}
