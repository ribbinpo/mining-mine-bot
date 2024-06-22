package repositories

import (
	"github.com/ribbinpo/mining-mine-bot/domain"
	"github.com/ribbinpo/mining-mine-bot/pkg/utils"
	"gorm.io/gorm"
)

type PriceTokenRepo struct {
	DB *gorm.DB
}

func NewPriceTokenRepository(db *gorm.DB) domain.PriceTokenRepository {
	return &PriceTokenRepo{DB: db}
}

func (p *PriceTokenRepo) RecordPriceToken(priceToken []*domain.PriceToken) error {
	result := p.DB.Create(priceToken)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (p *PriceTokenRepo) GetAll(pagination utils.Pagination, filter domain.PriceTokenFilter) ([]*domain.PriceToken, error) {
	var priceTokens []*domain.PriceToken
	offset := (pagination.Page - 1) * pagination.PerPage
	operate := p.DB
	if filter.CryptoCurrency != "" {
		operate = p.DB.Where("crypto_currency", filter.CryptoCurrency)
	}
	if filter.FiatAmounts != 0 {
		operate = operate.Where("amount_fiat_selected", filter.FiatAmounts)
	}
	result := operate.Limit(pagination.PerPage).Offset(offset).Find(&priceTokens)
	if result.Error != nil {
		return nil, result.Error
	}
	return priceTokens, nil
}
