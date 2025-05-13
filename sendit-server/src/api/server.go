package server

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"sendit.server/src/api/routes"
)

type Server struct{}

func BuildServer(ctx *gin.Context) *gin.Engine {
	router := gin.Default()

	emailRouter := routes.BuildEmailRouter(ctx)

	notifications := router.Group("/notifications")
	{
		emailRouter.GetRoutes(notifications)
	}

	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
