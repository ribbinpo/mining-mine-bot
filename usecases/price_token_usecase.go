package usecases

import (
	"github.com/ribbinpo/mining-mine-bot/domain"
	"github.com/ribbinpo/mining-mine-bot/pkg/utils"
)

type PriceTokenUseCase struct {
	repo domain.PriceTokenRepository
}

func NewPriceTokenUseCase(repo domain.PriceTokenRepository) domain.PriceTokenUsecase {
	return &PriceTokenUseCase{repo: repo}
}

func (p *PriceTokenUseCase) GetAll(pagination utils.Pagination, currency string) ([]*domain.PriceToken, error) {
	return p.repo.GetAll(pagination, currency)
}
