package routers

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Route(app *fiber.App, dbClient *gorm.DB) {
	api := app.Group("/api")

	NewTemplateController(api.Group("/template"), dbClient)
}
