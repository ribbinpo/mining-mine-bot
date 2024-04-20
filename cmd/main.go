package main

import (
	"fmt"

	"github.com/ribbinpo/mining-mine-bot/config"
	"github.com/ribbinpo/mining-mine-bot/domain"
	"github.com/ribbinpo/mining-mine-bot/pkg/database"
	"github.com/ribbinpo/mining-mine-bot/pkg/httpClient"
	"github.com/ribbinpo/mining-mine-bot/repositories"
	"github.com/ribbinpo/mining-mine-bot/usecases"
	"github.com/robfig/cron/v3"
)

func main() {
	cfg := config.NewConfig("./.env.development")
	client := &httpClient.DefaultClient{}

	c := cron.New(cron.WithLocation(cfg.App.TimeZone))
	defer c.Stop()

	dbClient := database.PostgresDBConnection(cfg.Db)
	dbClientSettings, _ := dbClient.DB()
	defer dbClientSettings.Close()

	// Migrate the schema
	dbClient.AutoMigrate(&domain.PriceToken{})

	// app := fiber.New()

	// app.Get("/health-check", func(c *fiber.Ctx) error {
	// 	return c.SendString("Hello, World!")
	// })

	// routers.Route(app, dbClient)

	c.AddFunc("@daily", func() {
		if err := usecases.NewP2PBinanceService(repositories.NewP2PBinanceRepository(client), repositories.NewPriceTokenRepository(dbClient)).RecordP2PBinanceData("https://p2p.binance.com/bapi/c2c/v2/friendly/c2c/adv/search"); err != nil {
			panic(err)
		}
		fmt.Println("Cron job executed")
	})

	c.Start()

	// u can remove this line if open fiber
	select {}
}
