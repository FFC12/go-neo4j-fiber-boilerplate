package utils

import (
	"os"

	"github.com/joho/godotenv"
)

var (
	DB_URI         string
	DB_NAME        string
	DB_USER        string
	DB_PASSWORD    string
	JWT_SECRET_KEY string
)

func InitConfig() {
	err := godotenv.Load()
	if err != nil {
		ErrorLogger.Panic("Error loading .env file")
	}

	DB_NAME = os.Getenv("DATABASE_NAME")
	DB_URI = os.Getenv("DATABASE_URI")
	DB_USER = os.Getenv("DATABASE_USER")
	DB_PASSWORD = os.Getenv("DATABASE_PASS")
	JWT_SECRET_KEY = os.Getenv("JWT_SECRET_KEY")
}
