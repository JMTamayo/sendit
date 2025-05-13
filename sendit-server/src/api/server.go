package server

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"sendit.server/src/api/routes"
	"sendit.server/src/models"
)

type Server struct{}

func BuildServer(ctx *gin.Context) (*gin.Engine, *models.Error) {
	router := gin.Default()

	emailRouter, err := routes.BuildEmailRouter(ctx)
	if err != nil {
		return nil, err
	}

	notifications := router.Group("/notifications")
	{
		emailRouter.GetRoutes(notifications)
	}

	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router, nil
}
