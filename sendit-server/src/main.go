package main

import (
	"sendit.server/src/api"
	"sendit.server/src/config"
	"sendit.server/src/docs"

	"github.com/gin-gonic/gin"
)

func main() {
	docs.SwaggerInfo.Title = "Sendit"
	docs.SwaggerInfo.Version = "1.0.0"

	ctx := gin.Context{}

	server := server.BuildServer(&ctx)

	server.Run(config.Conf.GetServiceAddress())
}
