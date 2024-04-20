package repositories

import (
	"github.com/ribbinpo/mining-mine-bot/domain"
	"gorm.io/gorm"
)

type PriceTokenRepo struct {
	DB *gorm.DB
}

func NewPriceTokenRepository(db *gorm.DB) domain.PriceTokenRepository {
	return &PriceTokenRepo{DB: db}
}

func (p *PriceTokenRepo) RecordPriceToken(priceToken []*domain.PriceToken) error {
	result := p.DB.Create(&priceToken)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
