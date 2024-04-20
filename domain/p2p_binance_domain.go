package domain

import "encoding/json"

type Advertiser struct {
	UserNo             string   `json:"userNo"`
	RealName           string   `json:"realName"`
	NickName           string   `json:"nickName"`
	Margin             string   `json:"margin"`
	MarginUnit         string   `json:"marginUnit"`
	OrderCount         string   `json:"orderCount"`
	MonthOrderCount    int      `json:"monthOrderCount"`
	MonthFinishRate    float64  `json:"monthFinishRate"`
	PositiveRate       float64  `json:"positiveRate"`
	AdvConfirmTime     string   `json:"advConfirmTime"`
	Email              string   `json:"email"`
	RegistrationTime   string   `json:"registrationTime"`
	Mobile             string   `json:"mobile"`
	UserType           string   `json:"userType"`
	TagIconUrls        []string `json:"tagIconUrls"`
	UserGrade          int      `json:"userGrade"`
	UserIdentity       string   `json:"userIdentity"`
	ProMerchant        string   `json:"proMerchant"`
	Badges             []string `json:"badges"`
	IsBlocked          string   `json:"isBlocked"`
	ActiveTimeInSecond int      `json:"activeTimeInSecond"`
}

type TradeMethod struct {
	PayId                string `json:"payId"`
	PayMethodId          string `json:"payMethodId"`
	PayType              string `json:"payType"`
	PayAccount           string `json:"payAccount"`
	PayBank              string `json:"payBank"`
	PaySubBank           string `json:"paySubBank"`
	Identifier           string `json:"identifier"`
	IconUrlColor         string `json:"iconUrlColor"`
	TradeMethodName      string `json:"tradeMethodName"`
	TradeMethodShortName string `json:"tradeMethodShortName"`
	TradeMethodBgColor   string `json:"tradeMethodBgColor"`
}

type Adv struct {
	AdvNo                           string        `json:"advNo"`
	Classify                        string        `json:"classify"`
	TradeType                       string        `json:"tradeType"`
	Asset                           string        `json:"asset"`
	FiatUnit                        string        `json:"fiatUnit"`
	AdvStatus                       string        `json:"advStatus"`
	PriceType                       string        `json:"priceType"`
	PriceFloatingRatio              string        `json:"priceFloatingRatio"`
	RateFloatingRatio               string        `json:"rateFloatingRatio"`
	CurrencyRate                    string        `json:"currencyRate"`
	Price                           string        `json:"price"`
	InitAmount                      string        `json:"initAmount"`
	SurplusAmount                   string        `json:"surplusAmount"`
	TradableQuantity                string        `json:"tradableQuantity"`
	AmountAfterEditing              string        `json:"amountAfterEditing"`
	MaxSingleTransAmount            string        `json:"maxSingleTransAmount"`
	MinSingleTransAmount            string        `json:"minSingleTransAmount"`
	BuyerKycLimit                   string        `json:"buyerKycLimit"`
	BuyerRegDaysLimit               string        `json:"buyerRegDaysLimit"`
	BuyerBtcPositionLimit           string        `json:"buyerBtcPositionLimit"`
	Remarks                         string        `json:"remarks"`
	AutoReplyMsg                    string        `json:"autoReplyMsg"`
	PayTimeLimit                    int           `json:"payTimeLimit"`
	TradeMethods                    []TradeMethod `json:"tradeMethods"`
	UserTradeCountFilterTime        string        `json:"userTradeCountFilterTime"`
	UserBuyTradeCountMin            string        `json:"userBuyTradeCountMin"`
	UserBuyTradeCountMax            string        `json:"userBuyTradeCountMax"`
	UserSellTradeCountMin           string        `json:"userSellTradeCountMin"`
	UserSellTradeCountMax           string        `json:"userSellTradeCountMax"`
	UserAllTradeCountMin            string        `json:"userAllTradeCountMin"`
	UserAllTradeCountMax            string        `json:"userAllTradeCountMax"`
	UserTradeCompleteRateFilterTime string        `json:"userTradeCompleteRateFilterTime"`
	UserTradeCompleteCountMin       string        `json:"userTradeCompleteCountMin"`
	UserTradeCompleteRateMin        string        `json:"userTradeCompleteRateMin"`
	UserTradeVolumeFilterTime       string        `json:"userTradeVolumeFilterTime"`
	UserTradeType                   string        `json:"userTradeType"`
	UserTradeVolumeMin              string        `json:"userTradeVolumeMin"`
	UserTradeVolumeMax              string        `json:"userTradeVolumeMax"`
	UserTradeVolumeAsset            string        `json:"userTradeVolumeAsset"`
	CreateTime                      string        `json:"createTime"`
	AdvUpdateTime                   string        `json:"advUpdateTime"`
	FiatVo                          string        `json:"fiatVo"`
	AssetVo                         string        `json:"assetVo"`
	AdvVisibleRet                   string        `json:"advVisibleRet"`
	TakerAdditionalKycRequired      int           `json:"takerAdditionalKycRequired"`
	InventoryType                   string        `json:"inventoryType"`
	OfflineReason                   string        `json:"offlineReason"`
	AssetLogo                       string        `json:"assetLogo"`
	AssetScale                      int           `json:"assetScale"`
	FiatScale                       int           `json:"fiatScale"`
	PriceScale                      int           `json:"priceScale"`
	FiatSymbol                      string        `json:"fiatSymbol"`
	IsTradable                      bool          `json:"isTradable"`
	DynamicMaxSingleTransAmount     string        `json:"dynamicMaxSingleTransAmount"`
	MinSingleTransQuantity          string        `json:"minSingleTransQuantity"`
	MaxSingleTransQuantity          string        `json:"maxSingleTransQuantity"`
	DynamicMaxSingleTransQuantity   string        `json:"dynamicMaxSingleTransQuantity"`
	CommissionRate                  string        `json:"commissionRate"`
	TakerCommissionRate             string        `json:"takerCommissionRate"`
	TradeMethodCommissionRates      []string      `json:"tradeMethodCommissionRates"`
	LaunchCountry                   string        `json:"launchCountry"`
	AbnormalStatusList              string        `json:"abnormalStatusList"`
	CloseReason                     string        `json:"closeReason"`
	StoreInformation                string        `json:"storeInformation"`
	AllowTradeMerchant              string        `json:"allowTradeMerchant"`
}

type P2PBinanceResponse struct {
	Code          string `json:"code"`
	Message       string `json:"message"`
	MessageDetail string `json:"messageDetail"`
	Data          []struct {
		Adv        Adv        `json:"adv"`
		Advertiser Advertiser `json:"advertiser"`
	} `json:"data"`
	Total   int  `json:"total"`
	Success bool `json:"success"`
}

type P2PBinanceDataPayload struct {
	AdditionalKycVerifyFilter int
	Asset                     string
	Classifies                []string
	Countries                 []string
	Fiat                      string
	FilterType                string
	Page                      int
	PayTypes                  []string
	ProMerchantAds            bool
	PublisherType             *string
	Rows                      int
	ShieldMerchantAds         bool
	TradeType                 string
	TransAmount               int
}

func (p P2PBinanceDataPayload) Encode() ([]byte, error) {
	body, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}
	return body, nil
}

// Repository
type P2PBinanceRepository interface {
	GetP2PBinanceData(url string, body []byte) (*P2PBinanceResponse, error)
}

// UseCase
type P2PBinanceUseCase interface {
	RecordP2PBinanceData(url string) error
}
