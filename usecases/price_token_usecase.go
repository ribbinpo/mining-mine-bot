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

func (p *PriceTokenUseCase) GetDiffPrice(filter domain.PriceTokenUseCaseGetDiffPriceFilter) (float64, error) {
	filter1 := domain.PriceTokenFilter{
		CryptoCurrency: filter.CryptoCurrency1,
		FiatAmounts:    filter.FiatAmounts,
	}
	filter2 := domain.PriceTokenFilter{
		CryptoCurrency: filter.CryptoCurrency2,
		FiatAmounts:    filter.FiatAmounts,
	}
	token1, err := p.repo.GetPriceTokenLastest(filter1)
	if err != nil {
		return 0, err
	}
	token2, err := p.repo.GetPriceTokenLastest(filter2)
	if err != nil {
		return 0, err
	}
	return token1.Price - token2.Price, nil
}
