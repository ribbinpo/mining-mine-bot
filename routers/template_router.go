package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ribbinpo/mining-mine-bot/controllers"
	"github.com/ribbinpo/mining-mine-bot/repositories"
	"github.com/ribbinpo/mining-mine-bot/usecases"
	"gorm.io/gorm"
)

func NewTemplateController(route fiber.Router, dbClient *gorm.DB) {

	controller := controllers.NewTemplateController(usecases.NewTemplateService(repositories.NewTemplateRepository(dbClient)))

	route.Get("/", controller.GetAll)
}
