package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Configuration struct {
	Db_Host     string
	Db_Port     int
	Db_Username string
	Db_Password string
	Db_Name     string
}

func LoadEnvVariables() (Configuration, error) {
	//COMMENT THIS IF U RUNNING THROUGH DOCKER
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("MASUK GODOT ENV LOAD ERROR")
		log.Fatal(err)
	}

	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		fmt.Println(err)
		fmt.Println("MASUK SINIII")
		fmt.Println("ERROR DI CONFIG", os.Getenv("DB_HOST"))
		return Configuration{}, err
	}

	return Configuration{
		Db_Host:     os.Getenv("DB_HOST"),
		Db_Port:     port,
		Db_Username: os.Getenv("DB_USERNAME"),
		Db_Password: os.Getenv("DB_PASSWORD"),
		Db_Name:     os.Getenv("DB_NAME"),
	}, nil
}
