package controllers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/ribbinpo/mining-mine-bot/domain"
	"github.com/ribbinpo/mining-mine-bot/pkg/utils"
)

type IPriceTokenController interface {
	GetAll(ctx *fiber.Ctx) error
	GetPriceTokenDescribe(ctx *fiber.Ctx) error
	GetPriceDiff(ctx *fiber.Ctx) error
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
	orderType := ctx.Query("type")

	var filter domain.PriceTokenFilter
	pagination := utils.Pagination{Page: page, PerPage: limit}
	if startDate == "" || endDate == "" {
		filter = domain.PriceTokenFilter{CryptoCurrency: currency, FiatAmounts: fiatAmounts, OrderType: orderType}
	} else {
		// layout := "Jan _2 15:04:05"
		parseStartDate, err := time.Parse(time.RFC3339, startDate)
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
		}
		parseEndDate, err := time.Parse(time.RFC3339, endDate)
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
		}
		filter = domain.PriceTokenFilter{CryptoCurrency: currency, FiatAmounts: fiatAmounts, StartDate: parseStartDate, EndDate: parseEndDate, OrderType: orderType}
	}

	result, err := p.PriceTokenService.GetAll(pagination, filter)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return ctx.Status(fiber.StatusOK).JSON(result)
}

func (p *PriceTokenController) GetPriceTokenDescribe(ctx *fiber.Ctx) error {
	currency := ctx.Query("currency")
	fiatAmounts := ctx.QueryInt("fiat_amounts", 0)
	orderType := ctx.Query("type")

	filter := domain.PriceTokenFilter{CryptoCurrency: currency, FiatAmounts: fiatAmounts, OrderType: orderType}
	result, err := p.PriceTokenService.GetPriceTokenDescribe(filter)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return ctx.Status(fiber.StatusOK).JSON(result)
}

func (p *PriceTokenController) GetPriceDiff(ctx *fiber.Ctx) error {
	currency1 := ctx.Query("currency1")
	currency2 := ctx.Query("currency2")
	fiatAmounts := ctx.QueryInt("fiat_amounts", 0)

	filter := domain.PriceTokenUseCaseGetDiffPriceFilter{CryptoCurrency1: currency1, CryptoCurrency2: currency2, FiatAmounts: fiatAmounts}
	result, err := p.PriceTokenService.GetDiffPrice(filter)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return ctx.Status(fiber.StatusOK).JSON(result)
}
