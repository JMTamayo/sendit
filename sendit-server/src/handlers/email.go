package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"sendit.server/src/config"
	"sendit.server/src/models"
	"sendit.server/src/services"
)

type EmailHandler struct {
	redisService *services.RedisService
}

func NewEmailHandler(ctx context.Context) *EmailHandler {
	return &EmailHandler{
		redisService: services.NewRedisService(ctx),
	}
}

// Send email service documentation
// @Summary Send an email
// @Description Send an email to a recipient with a specific subject and body. This endpoint produces a stream event that is processed by the email service.
// @Tags Notifications
// @Accept json
// @Produce json
// @Param email body models.Email true "Email"
// @Success 200 {object} models.SendEmailResponse
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /notifications/email [post]
func (h *EmailHandler) SendEmail(c *gin.Context) {
	config.Log.Info("Starting service to send email")

	var email models.Email
	if err := c.ShouldBindJSON(&email); err != nil {
		c.JSON(http.StatusBadRequest, models.NewError(
			fmt.Sprintf("Invalid request body: %v", err),
		))
		return
	}

	config.Log.Debug(fmt.Sprintf("Request body: %v", email))

	validate := validator.New()
	if err := validate.Struct(email); err != nil {
		c.JSON(http.StatusBadRequest, models.NewError(
			fmt.Sprintf("Invalid request body: %v", err),
		))
		return
	}

	data, marErr := json.Marshal(email)
	if marErr != nil {
		c.JSON(http.StatusInternalServerError, models.NewError(
			fmt.Sprintf("Error marshaling request body: %v", marErr),
		))
		return
	}

	message := map[string]string{"data": string(data)}

	id, err := h.redisService.Produce(c.Request.Context(), message)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.NewError(
			fmt.Sprintf("Error producing message to Redis: %v", err),
		))
		return
	}

	response := models.SendEmailResponse{
		Success: true,
		Details: "The email request has been received and it will be processed shortly",
		Email:   email,
		Id:      *id,
	}

	c.JSON(http.StatusOK, response)
}
