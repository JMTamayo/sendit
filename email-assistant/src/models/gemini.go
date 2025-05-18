package models

type GeminiResponse struct {
	IsValid bool   `json:"is_valid"`
	Details string `json:"details"`
}
