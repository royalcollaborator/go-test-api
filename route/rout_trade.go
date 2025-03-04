package route

import (
	"time"

	"github.com/royalcollaborator/go-test-api/api/controller"
	"github.com/royalcollaborator/go-test-api/bootstrap"
	"github.com/royalcollaborator/go-test-api/domain"

	"github.com/gin-gonic/gin"
)

func NewTradeApiRouter(env *bootstrap.Env, timeout time.Duration, group *gin.RouterGroup) {
	tradeService := &domain.TradeService{}
	sc := controller.TradeInfoController{
		TradeUsecase: tradeService,
		Env:          env,
	}
	group.GET("/ltp", sc.LastTradePrice)
}
