package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/ribbinpo/mining-mine-bot/config"
	"github.com/ribbinpo/mining-mine-bot/pkg/database"
	"github.com/ribbinpo/mining-mine-bot/routers"
)

func main() {
	cfg := config.NewConfig("./.env.development")

	dbClient := database.PostgresDBConnection(cfg.Db)
	dbClientSettings, _ := dbClient.DB()
	defer dbClientSettings.Close()

	app := fiber.New()

	app.Get("/health-check", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	routers.Route(app, dbClient)

	log.Fatal(app.Listen(cfg.App.Port))
}
