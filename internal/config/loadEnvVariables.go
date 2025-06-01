package config

import (
	"github.com/lpernett/godotenv"
	"log"
)

func LoadEnvVariables() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
