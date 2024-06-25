package initializers

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// LoadEnv loads the environment variables from the .env file.
func LoadEnv() {
	appEnv := os.Getenv("GO_ENVIRONMENT")
	if appEnv == "" {
		appEnv = "development"
	}

	if appEnv != "development" {
		return
	}

	if err := godotenv.Load(); err != nil {
		log.Fatal("Can't load .env file: ", err)
	}
}
