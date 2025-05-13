package routes

import (
	"github.com/gin-gonic/gin"

	"sendit.server/src/handlers"
)

type EmailRouter struct {
	EmailHandler *handlers.EmailHandler
}

func BuildEmailRouter(ctx *gin.Context) *EmailRouter {
	return &EmailRouter{
		EmailHandler: handlers.NewEmailHandler(ctx),
	}
}

func (r *EmailRouter) GetRoutes(routerGroup *gin.RouterGroup) {
	routerGroup.POST("/email", r.EmailHandler.SendEmail)
}
