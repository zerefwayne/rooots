package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnvVariables() {
	if IsEnvironmentHeroku() {
		return
	}

	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
}

func IsEnvironmentHeroku() bool {
	return os.Getenv("ENVIRONMENT") == "heroku"
}
