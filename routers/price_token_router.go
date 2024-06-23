package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ribbinpo/mining-mine-bot/controllers"
	"github.com/ribbinpo/mining-mine-bot/repositories"
	"github.com/ribbinpo/mining-mine-bot/usecases"
	"gorm.io/gorm"
)

func NewPriceTokenRouter(route fiber.Router, dbClient *gorm.DB) {

	controller := controllers.NewPriceTokenController(usecases.NewPriceTokenUseCase(repositories.NewPriceTokenRepository(dbClient)))

	route.Get("/", controller.GetAll)
	route.Get("/describe", controller.GetPriceTokenDescribe)
}
