package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"email.assistant/src/api/routes"
	"email.assistant/src/models"
)

type Server struct{}

func BuildServer(ctx *gin.Context) (*gin.Engine, *models.Error) {
	router := gin.Default()

	// Configure CORS
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	config.AllowMethods = []string{"POST", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Authorization"}
	config.AllowCredentials = true
	router.Use(cors.New(config))

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
