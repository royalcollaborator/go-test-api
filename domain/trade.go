package domain

//go:generate mockgen -destination mock/trade_mock.go ./ TradeUsecase

type Trade struct {
	Pair   string  `json:"pair"`
	Amount float64 `json:"amount"`
}

type LastTradeReponse struct {
	LTP []Trade `json:"ltp"`
}

type TradeUsecase interface {
	GetLastTradePrice() (statusCode int, body map[string]map[string]float64, err error)
}
