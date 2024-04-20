package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	App *App
	Db  *Db
}

func NewConfig(pathEnv string) *Config {
	err := godotenv.Load(pathEnv)
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	tz, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		log.Fatal("Error loading timezone")
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSLMODE"))
	url := fmt.Sprintf("%s:%s", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"))

	return &Config{
		App: &App{
			Port:     os.Getenv("PORT"),
			TimeZone: tz,
		},
		Db: &Db{
			Dsn: dsn,
			Url: url,
		},
	}
}

type App struct {
	Port     string
	TimeZone *time.Location
}

type Db struct {
	Url string
	Dsn string
}
