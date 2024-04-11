package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ribbinpo/mining-mine-bot/domain"
)

type TemplateController struct {
	TemplateService domain.TemplateUseCase
}

func NewTemplateController(service domain.TemplateUseCase) *TemplateController {
	return &TemplateController{TemplateService: service}
}

func (t *TemplateController) GetAll(ctx *fiber.Ctx) error {
	result, err := t.TemplateService.GetAll()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}
	return ctx.Status(fiber.StatusOK).JSON(result)
}
