package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	PORT           string
	FILE_PATH      string
	SECRET         string
	KEY_API        string
	CACHE_INTERVAL int
)

func ConfigInit() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Getting environment variables for config Failed")
	}

	PORT = os.Getenv("PORT")
	FILE_PATH = os.Getenv("FILE_PATH")
	SECRET = os.Getenv("SECRET")
	KEY_API = os.Getenv("KEY")
	CACHE_INTERVAL, _ = strconv.Atoi(os.Getenv("CACHE_DURATION"))

}
