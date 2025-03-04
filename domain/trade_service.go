package domain

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type TradeService struct{}

func (t *TradeService) GetLastTradePrice() (statusCode int, body map[string]map[string]float64, err error) {
	url := "https://api.coingecko.com/api/v3/simple/price?ids=bitcoin&vs_currencies=usd,chf,eur"
	lastVal, err := http.Get(url)
	if err != nil {
		return http.StatusBadGateway, nil, err
	}
	defer lastVal.Body.Close()

	statusCode = lastVal.StatusCode
	data, err := ioutil.ReadAll(lastVal.Body)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	var result map[string]map[string]float64

	err = json.Unmarshal(data, &result)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, result, nil
}
