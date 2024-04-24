package domain

import (
	"github.com/ribbinpo/mining-mine-bot/pkg/utils"
	"gorm.io/gorm"
)

type PriceToken struct {
	gorm.Model
	Price              float64 `json:"price"`
	CryptoCurrency     string  `json:"crypto_currency"`
	FiatCurrency       string  `json:"fiat_currency"`
	AmountFiatSelected uint    `json:"amount_fiat_selected"`
}

const (
	USDT = "USDT"
	BNB  = "BNB"
	ETH  = "ETH"
	BTC  = "BTC"
)

type PriceTokenRepository interface {
	GetAll(pagination utils.Pagination, currency string) ([]*PriceToken, error)
	RecordPriceToken(priceToken []*PriceToken) error
}

type PriceTokenUsecase interface {
	GetAll(pagination utils.Pagination, currency string) ([]*PriceToken, error)
}
