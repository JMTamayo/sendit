package main

import (
	"email.assistant/src/api"
	"email.assistant/src/config"
	"email.assistant/src/docs"

	"github.com/gin-gonic/gin"
)

func main() {
	docs.SwaggerInfo.Title = "Sendit - Email Assistant"
	docs.SwaggerInfo.Version = "1.0.1"

	ctx := gin.Context{}

	server, err := server.BuildServer(&ctx)
	if err != nil {
		panic(err)
	}

	server.Run(config.Conf.GetServiceAddress())
}
