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

type PriceTokenRepository interface {
	GetAll(pagination utils.Pagination, filter PriceTokenFilter) ([]*PriceToken, error)
	GetAvgPrice(filter PriceTokenFilter) (float64, error)
	GetPriceTokenDescribe(filter PriceTokenFilter) (*PriceTokenRepositoryDescribe, error)
	RecordPriceToken(priceToken []*PriceToken) error
}

type PriceTokenUsecase interface {
	GetAll(pagination utils.Pagination, filter PriceTokenFilter) (*PriceTokenUseCaseGetAllResponse, error)
	GetPriceTokenDescribe(filter PriceTokenFilter) (*PriceTokenRepositoryDescribe, error)
}
