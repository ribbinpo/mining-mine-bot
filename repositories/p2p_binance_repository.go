package repositories

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/ribbinpo/mining-mine-bot/domain"
	"github.com/ribbinpo/mining-mine-bot/pkg/httpClient"
)

type p2PBinanceRepo struct {
	HttpClient *httpClient.DefaultClient
}

func NewP2PBinanceRepository(HttpClient *httpClient.DefaultClient) domain.P2PBinanceRepository {
	return &p2PBinanceRepo{HttpClient: HttpClient}
}

func (p *p2PBinanceRepo) GetP2PBinanceData(url string, body []byte) (*domain.P2PBinanceResponse, error) {
	response, error := p.HttpClient.Post(url, body)

	if error != nil {
		return nil, fmt.Errorf("error while sending request: %v", error)
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("error while reading response body: %v", err)
	}

	var result domain.P2PBinanceResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("error while unmarshalling JSON: %v", err)
	}

	return &result, nil
}
