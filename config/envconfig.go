package config

import (
	"os"
)

// Config func to get env value from key ---
func EnvConfig(key string) string {

	// Dockerfile added, so godotenv module is not required
	// it's for local setup
	// load .env file
	// err := godotenv.Load(".env")
	// if err != nil {
	// 	log.Println("Error loading .env file")
	// }
	return os.Getenv(key)
}
