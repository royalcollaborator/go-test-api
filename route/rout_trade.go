package route

import (
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/royalcollaborator/go-test-api/api/controller"
	"github.com/royalcollaborator/go-test-api/api/middleware"
	"github.com/royalcollaborator/go-test-api/bootstrap"
	"github.com/royalcollaborator/go-test-api/domain"

	"github.com/gin-gonic/gin"
)

func NewTradeApiRouter(env *bootstrap.Env, timeout time.Duration, rdb *redis.Client, group *gin.RouterGroup) {
	tradeService := &domain.TradeService{}
	sc := controller.TradeInfoController{
		TradeUsecase: tradeService,
		Env:          env,
	}
	ttl := time.Duration(env.CacheTTL) * time.Second
	group.GET("/ltp", middleware.CacheMiddleware(rdb, ttl), sc.LastTradePrice)
}
