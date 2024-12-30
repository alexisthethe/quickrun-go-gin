package initializers

import (
	"github.com/joho/godotenv"
)

func LoadEnvVariables() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}