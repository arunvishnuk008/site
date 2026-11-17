package initializers

import (
	"github.com/joho/godotenv"
	"log"
)

func LoadEnvVariables() {
	err := godotenv.Load()
	if err != nil {
		log.Print("Error loading .env file, falling back to system env assuming those are set.")
	}
}
