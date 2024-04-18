package usecases

import (
	"fmt"
	"sync"

	"github.com/ribbinpo/mining-mine-bot/domain"
)

type P2PBinanceService struct {
	repo domain.P2PBinanceRepository
}

func NewP2PBinanceService(repo domain.P2PBinanceRepository) domain.P2PBinanceUseCase {
	return &P2PBinanceService{repo: repo}
}

func (s *P2PBinanceService) RecordP2PBinanceData(url string) error {

	fiatAmount := 1000
	assets := []string{"USDT", "BTC", "ETH"}
	var wg sync.WaitGroup

	wg.Add(len(assets))

	for _, asset := range assets {
		_asset := asset
		go func() {
			payload := domain.P2PBinanceDataPayload{
				AdditionalKycVerifyFilter: 0,
				Asset:                     _asset,
				Classifies:                []string{"mass", "profession"},
				Countries:                 []string{},
				Fiat:                      "THB",
				FilterType:                "all",
				Page:                      1,
				PayTypes:                  []string{},
				ProMerchantAds:            false,
				PublisherType:             nil,
				Rows:                      2,
				ShieldMerchantAds:         false,
				TradeType:                 "BUY",
				TransAmount:               fiatAmount,
			}
			body, err := payload.Encode()
			if err != nil {
				fmt.Println("Error marshalling JSON:", err)
				return
			}
			result, err := s.repo.GetP2PBinanceData(url, body)
			if err != nil {
				fmt.Println("Error getting data:", err)
				return
			}
			fmt.Println(_asset)
			for _, data := range result.Data {
				fmt.Println(data.Adv.Price)
			}
			fmt.Println(result.Success)
			wg.Done()
		}()
	}

	wg.Wait()

	// find the best price

	// record price

	return nil
}
