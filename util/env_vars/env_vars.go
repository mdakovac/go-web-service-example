package env_vars

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type EnvVariablesType struct {
	DATABASE_URL string
}

var EnvVariables EnvVariablesType

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	EnvVariables = EnvVariablesType{}

	EnvVariables.DATABASE_URL = os.Getenv("DATABASE_URL")
}
