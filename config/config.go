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

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
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

	// dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
	// 	os.Getenv("DB_HOST"),
	// 	os.Getenv("DB_USER"),
	// 	os.Getenv("DB_PASSWORD"),
	// 	os.Getenv("DB_NAME"),
	// 	os.Getenv("DB_PORT"),
	// 	os.Getenv("DB_SSLMODE"))
	// url := fmt.Sprintf("%s:%s", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"))

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		getEnv("DB_HOST", os.Getenv("DB_HOST")),
		getEnv("DB_USER", os.Getenv("DB_USER")),
		getEnv("DB_PASSWORD", os.Getenv("DB_PASSWORD")),
		getEnv("DB_NAME", os.Getenv("DB_NAME")),
		getEnv("DB_PORT", os.Getenv("DB_PORT")),
		getEnv("DB_SSLMODE", os.Getenv("DB_SSLMODE")))
	url := fmt.Sprintf("%s:%s", getEnv("DB_HOST", os.Getenv("DB_HOST")), getEnv("DB_PORT", os.Getenv("DB_PORT")))

	return &Config{
		App: &App{
			// Port:     ":" + os.Getenv("PORT"),
			Port:     ":" + getEnv("PORT", os.Getenv("PORT")),
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
