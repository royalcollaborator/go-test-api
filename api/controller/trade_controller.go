package controller

import (
	"net/http"

	"github.com/royalcollaborator/go-test-api/bootstrap"
	"github.com/royalcollaborator/go-test-api/domain"

	"github.com/gin-gonic/gin"
)

type TradeInfoController struct {
	TradeUsecase domain.TradeUsecase
	Env          *bootstrap.Env
}

func (sc *TradeInfoController) LastTradePrice(c *gin.Context) {
	statusCode, lastVal, err := sc.TradeUsecase.GetLastTradePrice()
	if err != nil {
		c.JSON(statusCode, domain.ErrorResponse{Message: err.Error()})
		return
	}

	if statusCode != http.StatusOK {
		c.JSON(statusCode, domain.ErrorResponse{Message: err.Error()})
		return
	}

	res := domain.LastTradeReponse{
		LTP: []domain.Trade{
			{Pair: "BTC/CHF", Amount: lastVal["bitcoin"]["chf"]},
			{Pair: "BTC/EUR", Amount: lastVal["bitcoin"]["eur"]},
			{Pair: "BTC/USD", Amount: lastVal["bitcoin"]["usd"]},
		},
	}
	c.JSON(statusCode, res)
}
