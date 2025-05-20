package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"email.assistant/src/api/routes"
	"email.assistant/src/config"
	"email.assistant/src/models"
)

type Server struct{}

func BuildServer(ctx *gin.Context) (*gin.Engine, *models.Error) {
	router := gin.Default()

	// Configure CORS
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = config.Conf.GetAllowedOrigins()
	corsConfig.AllowMethods = []string{"POST", "OPTIONS"}
	corsConfig.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Authorization"}
	corsConfig.AllowCredentials = true
	router.Use(cors.New(corsConfig))

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
