package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/ribbinpo/mining-mine-bot/config"
	"github.com/ribbinpo/mining-mine-bot/domain"
	"github.com/ribbinpo/mining-mine-bot/pkg/database"
	"github.com/ribbinpo/mining-mine-bot/pkg/httpClient"
	"github.com/ribbinpo/mining-mine-bot/repositories"
	"github.com/ribbinpo/mining-mine-bot/routers"
	"github.com/ribbinpo/mining-mine-bot/usecases"
	"github.com/robfig/cron/v3"
)

func main() {
	cfg := config.NewConfig("./.env")
	client := &httpClient.DefaultClient{}

	c := cron.New(cron.WithLocation(cfg.App.TimeZone))
	defer c.Stop()

	dbClient := database.PostgresDBConnection(cfg.Db)
	dbClientSettings, _ := dbClient.DB()
	defer dbClientSettings.Close()

	// Migrate the schema
	dbClient.AutoMigrate(&domain.PriceToken{})

	app := fiber.New()
	app.Use(cors.New())

	c.AddFunc("@daily", func() {
		p2pBinanceUsecase := usecases.NewP2PBinanceService(repositories.NewP2PBinanceRepository(client), repositories.NewPriceTokenRepository(dbClient), "https://p2p.binance.com/bapi/c2c/v2/friendly/c2c/adv/search")
		if err := p2pBinanceUsecase.RecordP2PBinanceDataForBuy(); err != nil {
			panic(err)
		}
		if err := p2pBinanceUsecase.RecordP2PBinanceDataForSell(); err != nil {
			panic(err)
		}
		fmt.Println("Cron job executed")
	})

	c.Start()

	app.Get("/health-check", func(c *fiber.Ctx) error {
		return c.SendString("Server is running:)")
	})

	routers.Route(app, dbClient)

	app.Listen(cfg.App.Port)

	// u can remove this line if open fiber
	// select {}
}
