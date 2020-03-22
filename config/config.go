package config

import (
	_ "github.com/joho/godotenv/autoload"
	"os"
)

type Config struct {
	AppPort    string
	DBPort     string
	DBHost     string
	DBPassword string
	DBUser     string
	DBName     string
}

func GetConfig() *Config {
	return &Config{
		AppPort:    os.Getenv("APPHOST"),
		DBPort:     os.Getenv("DBPORT"),
		DBHost:     os.Getenv("DBHOST"),
		DBPassword: os.Getenv("DBPASSWORD"),
		DBUser:     os.Getenv("DBUSER"),
		DBName:     os.Getenv("DBNAME"),
	}
}
