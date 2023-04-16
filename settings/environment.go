package settings

import (
	"log"

	"github.com/joho/godotenv"
)

func setEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

}
