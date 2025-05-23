package routes

import (
	"github.com/gin-gonic/gin"

	"email.assistant/src/handlers"
	"email.assistant/src/models"
)

type EmailRouter struct {
	EmailHandler *handlers.EmailHandler
}

func BuildEmailRouter(ctx *gin.Context) (*EmailRouter, *models.Error) {
	emailHandler, err := handlers.NewEmailHandler(ctx)
	if err != nil {
		return nil, err
	}

	return &EmailRouter{
		EmailHandler: emailHandler,
	}, nil
}

func (r *EmailRouter) GetRoutes(routerGroup *gin.RouterGroup) {
	routerGroup.POST("/email", r.EmailHandler.SendEmail)
}
