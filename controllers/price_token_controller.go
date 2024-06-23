package controllers

import (
	"time"

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
	fiatAmounts := ctx.QueryInt("fiat_amounts", 0)
	startDate := ctx.Query("start_date")
	endDate := ctx.Query("end_date")

	// layout := "Jan _2 15:04:05"
	parseStartDate, err := time.Parse(time.RFC3339, startDate)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	parseEndDate, err := time.Parse(time.RFC3339, endDate)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	filter := domain.PriceTokenFilter{CryptoCurrency: currency, FiatAmounts: fiatAmounts, StartDate: parseStartDate, EndDate: parseEndDate}
	pagination := utils.Pagination{Page: page, PerPage: limit}
	result, err := p.PriceTokenService.GetAll(pagination, filter)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return ctx.Status(fiber.StatusOK).JSON(result)
}
