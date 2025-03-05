package middleware

import (
	"bytes"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func CacheMiddleware(redisClient *redis.Client, ttl time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		cacheKey := "cache:" + c.Request.URL.RequestURI()

		cachedResponse, err := redisClient.Get(c, cacheKey).Result()
		if err == nil {
			c.Data(http.StatusOK, "application/json", []byte(cachedResponse))
			c.Abort()
			return
		}

		writer := &bodyWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = writer

		c.Next()

		redisClient.Set(c, cacheKey, writer.body.String(), ttl)
	}
}

type bodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}
