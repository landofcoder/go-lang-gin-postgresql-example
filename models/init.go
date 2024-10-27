package models

import (
	"github.com/joho/godotenv"
	"log"
)

func InitEnv() {

	err := godotenv.Load("app.env")

	if err != nil {
		log.Fatal("Error loading app.env file")
	}
}
