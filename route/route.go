package route

import (
	"time"

	"github.com/royalcollaborator/go-test-api/bootstrap"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func Setup(env *bootstrap.Env, timeout time.Duration, rdb *redis.Client, gin *gin.Engine) {
	apiRouter := gin.Group("/api/v1")
	NewTradeApiRouter(env, timeout, rdb, apiRouter)
}
