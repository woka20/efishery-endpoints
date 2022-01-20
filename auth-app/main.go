package main

import (
	"efishery-endpoints/auth-app/config"

	"github.com/kataras/iris/v12"
)

func main() {

	config.ConfigInit()

	apps := iris.New()

	apps.Handle()
}
