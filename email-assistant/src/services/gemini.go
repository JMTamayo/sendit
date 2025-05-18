package services

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"google.golang.org/genai"

	"email.assistant/src/config"
	"email.assistant/src/models"
)

type GeminiService struct {
	client *genai.Client
}

func (s *GeminiService) getClient() *genai.Client {
	return s.client
}

func NewGeminiService(ctx context.Context) (*GeminiService, *models.Error) {
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  config.Conf.GetGeminiAPIKey(),
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		return nil, models.NewError(fmt.Sprintf("Failed to create AI client: %v", err))
	}

	return &GeminiService{
		client: client,
	}, nil
}

func (g *GeminiService) getFullPrompt(email string) string {
	return `[CONTEXT]
	You are an email validation assistant. Your task is to analyze emails and validate them according to specific rules.

	[RULES]
	1. Language: The email must be written in English or Spanish. It can contain specific words in other languages, but the main content must be in English or Spanish.
	2. Content: The email must not contain offensive language about companies, people, governments, or animals.
	3. Semantics: The email must be semantically understandable.
	4. Spelling: The email must not contain spelling errors.
	5. Privacy: The email must not share credentials or personal information about any person, company, or government.
	6. Length: The email must not exceed 500 words.
	7. Structure: The email must include a greeting and a farewell.
	8. Tone: The email must be written in neutral language, without formalisms.

	[RESPONSE FORMAT]
	You must answer with a JSON object containing:
	- is_valid: boolean (true if all rules are met, false otherwise)
	- details: string (null if is_valid is true, or a brief explanation of validation findings if false)

	Don't answer me in markdown format, you must provide me with the JSON as a string.

	[EMAIL TO ANALYZE]
	` + email
}

func (s *GeminiService) VerifyEmail(ctx context.Context, email string) *models.Error {
	prompt := s.getFullPrompt(email)

	response, err := s.getClient().Models.GenerateContent(
		ctx,
		config.Conf.GetGeminiModel(),
		[]*genai.Content{
			{
				Parts: []*genai.Part{
					{
						Text: prompt,
					},
				},
			},
		},
		nil,
	)
	if err != nil {
		return models.NewError(fmt.Sprintf("Failed to generate content: %v", err))
	}

	responseContent := response.Text()
	responseContent = strings.ReplaceAll(responseContent, "```json\n", "")
	responseContent = strings.ReplaceAll(responseContent, "\n```", "")

	validation := models.GeminiResponse{}
	err = json.Unmarshal([]byte(responseContent), &validation)
	if err != nil {
		return models.NewError(fmt.Sprintf("Failed to unmarshal response from AI assistant: %v", err))
	}

	if !validation.IsValid {
		return models.NewError(validation.Details)
	}

	return nil
}
