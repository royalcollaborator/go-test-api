package domain

import (
	"encoding/json"
	"net/http"
)

type TradeService struct{}

type MarketData struct {
	CurrentPrice map[string]float64 `json:"current_price"`
	LastUpdated  string             `json:"last_updated"`
}

type CoinResponse struct {
	MarketData  MarketData `json:"market_data"`
	LastUpdated string     `json:"last_updated"`
}

func (t *TradeService) GetLastTradePrice() (statusCode int, body map[string]map[string]float64, err error) {
	url := "https://api.coingecko.com/api/v3/coins/bitcoin?localization=false&tickers=false&market_data=true&community_data=false&developer_data=false&sparkline=false"
	lastVal, err := http.Get(url)

	if err != nil {
		return http.StatusBadGateway, nil, err
	}
	defer lastVal.Body.Close()

	var result CoinResponse
	if err := json.NewDecoder(lastVal.Body).Decode(&result); err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, map[string]map[string]float64{
		"bitcoin": {
			"usd": result.MarketData.CurrentPrice["usd"],
			"chf": result.MarketData.CurrentPrice["chf"],
			"eur": result.MarketData.CurrentPrice["eur"],
		},
	}, nil
}
