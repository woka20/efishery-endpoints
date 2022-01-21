package main

import (
	"efishery-endpoints/fetch-app/cache"
	"efishery-endpoints/fetch-app/config"
	"efishery-endpoints/fetch-app/handler/commodity"
	"efishery-endpoints/fetch-app/handler/tokenization"
	rt "efishery-endpoints/fetch-app/routing"
	"fmt"
	"time"

	"github.com/kataras/iris/v12"
)

func main() {
	config.ConfigInit()
	apps := iris.New()

	cacheTool := cache.NewCache()
	cacheTimer := time.NewTicker(time.Duration(config.CACHE_INTERVAL) * time.Minute)
	go cacheTool.StartCache(cacheTimer)

	commoHandler := commodity.NewComodityHandler()

	apps.Get("/commodityList", rt.JWTIsExist, commoHandler.GetList)
	apps.Get("/aggregateData", rt.JWTIsExist, commoHandler.AggregateData)

	tokenHandler := tokenization.NewTokenHandler()
	apps.Get("/claims", rt.JWTIsExist, tokenHandler.GetClaims)

	listenPort := config.PORT
	fmt.Println("Server Ready")

	apps.Listen(":" + listenPort)

}
