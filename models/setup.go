package models

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB

func ConnectDatabase() {

	dsn := "host=localhost user=" + os.Getenv("POSTGRES_USER") + " password=" + os.Getenv("POSTGRES_PASSWORD") + " dbname=" + os.Getenv("POSTGRES_DB") + " port=" + os.Getenv("DB_POST") + " sslmode=" + os.Getenv("SSLMODE") + " TimeZone=" + os.Getenv("TIMEZONE")
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	err = database.AutoMigrate(&Contact{})
	if err != nil {
		return
	}

	DB = database
}
