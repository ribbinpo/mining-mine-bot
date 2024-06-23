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

func (p *PriceTokenUseCase) GetAll(pagination utils.Pagination, filter domain.PriceTokenFilter) (*domain.PriceTokenUseCaseGetAllResponse, error) {
	priceList, err := p.repo.GetAll(pagination, filter)
	if err != nil {
		return nil, err
	}
	avgPrice, err := p.repo.GetAvgPrice(filter)
	if err != nil {
		return nil, err
	}
	data := &domain.PriceTokenUseCaseGetAllResponse{
		AvgPrice:     avgPrice,
		LastestPrice: priceList[len(priceList)-1].Price,
		Data:         priceList,
	}

	return data, nil
}

func (p *PriceTokenUseCase) GetPriceTokenDescribe(filter domain.PriceTokenFilter) (*domain.PriceTokenRepositoryDescribe, error) {
	priceTokenDescribe, err := p.repo.GetPriceTokenDescribe(filter)
	if err != nil {
		return nil, err
	}
	return priceTokenDescribe, nil
}
