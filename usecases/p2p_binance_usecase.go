package usecases

import (
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/ribbinpo/mining-mine-bot/domain"
)

type P2PBinanceService struct {
	p2PBinanceRepo domain.P2PBinanceRepository
	priceTokenRepo domain.PriceTokenRepository
}

func NewP2PBinanceService(p2PBinanceRepo domain.P2PBinanceRepository, priceTokenRepo domain.PriceTokenRepository) domain.P2PBinanceUseCase {
	return &P2PBinanceService{p2PBinanceRepo: p2PBinanceRepo, priceTokenRepo: priceTokenRepo}
}

func (s *P2PBinanceService) RecordP2PBinanceData(url string) error {
	timeStart := time.Now()
	fiatAmounts := []int{1000, 10000, 100000}
	assets := []string{"USDT", "BTC", "ETH"}
	length := len(assets) * len(fiatAmounts)
	var wg sync.WaitGroup
	var mutex sync.Mutex
	// ch := make(chan domain.PriceToken, length)
	priceTokenLists := []*domain.PriceToken{}

	wg.Add(length)

	for _, asset := range assets {
		_asset := asset
		for _, fiatAmount := range fiatAmounts {
			_fiatAmount := fiatAmount
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
					Rows:                      1,
					ShieldMerchantAds:         false,
					TradeType:                 "BUY",
					TransAmount:               _fiatAmount,
				}
				body, err := payload.Encode()
				if err != nil {
					fmt.Println("Error marshalling JSON:", err)
					return
				}
				result, err := s.p2PBinanceRepo.GetP2PBinanceData(url, body)
				if err != nil {
					fmt.Println("Error getting data:", err)
					return
				}
				price := float64(-1)
				if len(result.Data) > 0 {
					fmt.Printf("Asset: %s, FiatAmount: %d, Price: %s\n", _asset, _fiatAmount, result.Data[0].Adv.Price)
					_price, err := strconv.ParseFloat(result.Data[0].Adv.Price, 64)
					price = _price
					if err != nil {
						fmt.Println("Error parsing price")
					}
				}

				mutex.Lock()
				// ch <- domain.PriceToken{
				// 	Price:              price,
				// 	CryptoCurrency:     _asset,
				// 	FiatCurrency:       "THB",
				// 	AmountFiatSelected: uint(_fiatAmount),
				// }
				priceTokenLists = append(priceTokenLists, &domain.PriceToken{
					Price:              price,
					CryptoCurrency:     _asset,
					FiatCurrency:       "THB",
					AmountFiatSelected: uint(_fiatAmount),
				})
				mutex.Unlock()
				wg.Done()
			}()
		}
	}

	wg.Wait()

	if err := s.priceTokenRepo.RecordPriceToken(priceTokenLists); err != nil {
		fmt.Println("Error recording price token:", err)
		return err
	}

	fmt.Print("Time taken: ", time.Since(timeStart))

	return nil
}
