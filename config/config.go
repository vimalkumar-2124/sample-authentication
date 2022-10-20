package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config func to get env value from key ---
func Config(key string) string {
	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Error loading .env file")
	}
	// log.Println(".env file :", key)
	return os.Getenv(key)
}
