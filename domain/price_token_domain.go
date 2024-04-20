package domain

import "gorm.io/gorm"

type PriceToken struct {
	gorm.Model
	Price              float64 `json:"price"`
	CryptoCurrency     string  `json:"crypto_currency"`
	FiatCurrency       string  `json:"fiat_currency"`
	AmountFiatSelected uint    `json:"amount_fiat_selected"`
}

type PriceTokenRepository interface {
	RecordPriceToken(priceToken []*PriceToken) error
}
