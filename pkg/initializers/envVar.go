package initializers

import (
	"github.com/joho/godotenv"
	"log"
)

// LoadEnv loads the environment variables from the .env file.
func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Can't load .env file: ", err)
	}
}
