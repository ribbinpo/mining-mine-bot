package database

import (
	"context"
	"log"
	"time"

	"github.com/ribbinpo/mining-mine-bot/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func PostgresDBConnection(cfgDb *config.Db) *gorm.DB {
	_, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	db, err := gorm.Open(postgres.Open(cfgDb.Dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Connection to url %s failed: %v", cfgDb.Url, err)
	}

	if err := db.AutoMigrate(); err != nil {
		log.Fatal("AutoMigrate failed: %v", err)
	}

	// NOTE: if you want to close the connection, you can use the following code
	// dbInstance, _ := db.DB()
	// defer dbInstance.Close()

	return db
}
