package main

import (
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/royalcollaborator/go-test-api/bootstrap"
	route "github.com/royalcollaborator/go-test-api/route"

	"github.com/gin-gonic/gin"
)

func main() {
	app := bootstrap.App()
	env := app.Env
	timeout := time.Duration(env.ContextTimeout) * time.Second
	gin := gin.Default()
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", env.RedisHost, env.RedisPort),
		Password: env.RedisPassword,
		DB:       0,
	})
	defer rdb.Close()
	route.Setup(env, timeout, rdb, gin)
	gin.Run(env.ServerAddress)
}
