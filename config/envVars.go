package config

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvVariables() {
	err := godotenv.Load()
	fmt.Println("MASUK LOAD ENV VARIABLE NIH")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
