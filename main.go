package main

import (
	"time"

	"github.com/royalcollaborator/go-test-api/bootstrap"
	route "github.com/royalcollaborator/go-test-api/route"

	"github.com/gin-gonic/gin"
)

func main() {
	app := bootstrap.App()
	env := app.Env
	timeout := time.Duration(env.ContextTimeout) * time.Second
	gin := gin.Default()
	route.Setup(env, timeout, gin)
	gin.Run(env.ServerAddress)
}
