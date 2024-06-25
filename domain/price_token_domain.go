package domain

import (
	"time"

	"github.com/ribbinpo/mining-mine-bot/pkg/utils"
	"gorm.io/gorm"
)

type PriceToken struct {
	gorm.Model
	Price              float64 `json:"price"`
	CryptoCurrency     string  `json:"crypto_currency"`
	FiatCurrency       string  `json:"fiat_currency"`
	AmountFiatSelected uint    `json:"amount_fiat_selected"`
	Type               string  `json:"type"`
}

const (
	USDT = "USDT"
	BNB  = "BNB"
	ETH  = "ETH"
	BTC  = "BTC"
)

type PriceTokenFilter struct {
	CryptoCurrency string
	FiatAmounts    int
	StartDate      time.Time
	EndDate        time.Time
	OrderType      string
}

type PriceTokenUseCaseGetAllResponse struct {
	AvgPrice     float64       `json:"avg_price"`
	LastestPrice float64       `json:"lastest_price"`
	Data         []*PriceToken `json:"data"`
}

type PriceTokenRepositoryDescribe struct {
	AvgPrice     float64 `json:"avg_price"`
	LastestPrice float64 `json:"lastest_price"`
	MinPrice     float64 `json:"min_price"`
	MaxPrice     float64 `json:"max_price"`
}

type PriceTokenUseCaseGetDiffPriceFilter struct {
	CryptoCurrency1 string
	CryptoCurrency2 string
	FiatAmounts     int
}

type PriceTokenRepository interface {
	GetAll(pagination utils.Pagination, filter PriceTokenFilter) ([]*PriceToken, error)
	GetPriceTokenDescribe(filter PriceTokenFilter) (*PriceTokenRepositoryDescribe, error)
	GetPriceTokenLastest(filter PriceTokenFilter) (*PriceToken, error)
	RecordPriceToken(priceToken []*PriceToken) error
}

type PriceTokenUsecase interface {
	GetAll(pagination utils.Pagination, filter PriceTokenFilter) (*PriceTokenUseCaseGetAllResponse, error)
	GetPriceTokenDescribe(filter PriceTokenFilter) (*PriceTokenRepositoryDescribe, error)
	GetDiffPrice(filter PriceTokenUseCaseGetDiffPriceFilter) (float64, error)
}
