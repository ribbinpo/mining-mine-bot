package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ribbinpo/mining-mine-bot/domain"
	"github.com/ribbinpo/mining-mine-bot/pkg/utils"
)

type IPriceTokenController interface {
	GetAll(ctx *fiber.Ctx) error
}

type PriceTokenController struct {
	PriceTokenService domain.PriceTokenUsecase
}

func NewPriceTokenController(service domain.PriceTokenUsecase) IPriceTokenController {
	return &PriceTokenController{PriceTokenService: service}
}

func (p *PriceTokenController) GetAll(ctx *fiber.Ctx) error {
	page := ctx.QueryInt("page", 1)
	limit := ctx.QueryInt("limit", 10)
	currency := ctx.Query("currency")
	pagination := utils.Pagination{Page: page, PerPage: limit}
	result, err := p.PriceTokenService.GetAll(pagination, currency)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return ctx.Status(fiber.StatusOK).JSON(result)
}
