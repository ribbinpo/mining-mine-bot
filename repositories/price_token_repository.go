package repositories

import (
	"time"

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
	// offset := (pagination.Page - 1) * pagination.PerPage
	operate := p.DB
	if filter.CryptoCurrency != "" {
		operate = p.DB.Where("crypto_currency", filter.CryptoCurrency)
	}
	if filter.FiatAmounts != 0 {
		operate = operate.Where("amount_fiat_selected", filter.FiatAmounts)
	}
	if (filter.StartDate != time.Time{} && filter.EndDate != time.Time{}) {
		operate = operate.Where("created_at between ? AND ?", filter.StartDate, filter.EndDate)
	}
	result := operate.Find(&priceTokens)
	if result.Error != nil {
		return nil, result.Error
	}
	return priceTokens, nil
}

func (p *PriceTokenRepo) GetPriceTokenDescribe(filter domain.PriceTokenFilter) (*domain.PriceTokenRepositoryDescribe, error) {
	var priceTokenDescribe domain.PriceTokenRepositoryDescribe
	result := p.DB.Table("price_tokens").Select("AVG(price) as avg_price", "MIN(price) as min_price", "MAX(price) as max_price").Where("type", filter.OrderType).Where("crypto_currency = ?", filter.CryptoCurrency).Where("amount_fiat_selected", filter.FiatAmounts).Find(&priceTokenDescribe)
	if result.Error != nil {
		return nil, result.Error
	}
	latestPriceResult := p.DB.Table("price_tokens").Select("price").Where("type", filter.OrderType).Where("crypto_currency = ?", filter.CryptoCurrency).Where("amount_fiat_selected", filter.FiatAmounts).Order("created_at desc").Limit(1).Find(&priceTokenDescribe.LastestPrice)
	if latestPriceResult.Error != nil {
		return nil, latestPriceResult.Error
	}
	return &priceTokenDescribe, nil
}

func (p *PriceTokenRepo) GetPriceTokenLastest(filter domain.PriceTokenFilter) (*domain.PriceToken, error) {
	var priceToken domain.PriceToken
	result := p.DB.Table("price_tokens").Where("type", filter.OrderType).Where("crypto_currency = ?", filter.CryptoCurrency).Where("amount_fiat_selected", filter.FiatAmounts).Order("created_at desc").Limit(1).Find(&priceToken)
	if result.Error != nil {
		return nil, result.Error
	}
	return &priceToken, nil
}
