package models

type Email struct {
	Recipient string `json:"recipient" validate:"required,email"`
	Subject   string `json:"subject" validate:"required"`
	Body      string `json:"body" validate:"required"`
}

type SendEmailResponse struct {
	Success bool   `json:"success"`
	Details string `json:"details"`
	Email   Email  `json:"email"`
	Id      string `json:"id"`
}
