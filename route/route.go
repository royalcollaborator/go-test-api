package route

import (
	"time"

	"github.com/royalcollaborator/go-test-api/bootstrap"

	"github.com/gin-gonic/gin"
)

func Setup(env *bootstrap.Env, timeout time.Duration, gin *gin.Engine) {
	apiRouter := gin.Group("/api/v1")
	NewTradeApiRouter(env, timeout, apiRouter)
}
