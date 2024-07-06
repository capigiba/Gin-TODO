package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file in your project")
	}
}

func GetDBConnectionString() string {
	return os.Getenv("DB_CONNECTION_STRING")
}
