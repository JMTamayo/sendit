package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"email.assistant/src/config"
	"email.assistant/src/models"
	"email.assistant/src/services"
)

type EmailHandler struct {
	redisService  *services.RedisService
	geminiService *services.GeminiService
}

func (h *EmailHandler) getRedisService() *services.RedisService {
	return h.redisService
}

func (h *EmailHandler) getGeminiService() *services.GeminiService {
	return h.geminiService
}

func NewEmailHandler(ctx context.Context) (*EmailHandler, *models.Error) {
	redisService, err := services.NewRedisService(ctx)
	if err != nil {
		return nil, err
	}

	geminiService, err := services.NewGeminiService(ctx)
	if err != nil {
		return nil, err
	}

	return &EmailHandler{
		redisService:  redisService,
		geminiService: geminiService,
	}, nil
}

// Send email service documentation
// @Summary Send an email
// @Description Send an email to a recipient with a specific subject and body. This endpoint produces a stream event that is processed by the email service.
// @Tags Notifications
// @Accept json
// @Produce json
// @Param email body models.Email true "Email"
// @Success 201 {object} models.SendEmailResponse
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

	config.Log.Debug(fmt.Sprintf("Email to verify: %v", string(data)))

	validationErr := h.getGeminiService().VerifyEmail(c.Request.Context(), email.Body)
	if validationErr != nil {
		c.JSON(http.StatusBadRequest, models.NewError(
			fmt.Sprintf("Invalid email: %v", *validationErr),
		))
		return
	}

	config.Log.Debug("Email successfully verified")

	message := map[string]string{config.Conf.GetKeyNameData(): string(data)}

	id, err := h.getRedisService().Produce(c.Request.Context(), message)
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

	c.JSON(http.StatusCreated, response)
}
