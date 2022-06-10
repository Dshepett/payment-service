package models

import (
	"net/mail"
	"time"
)

var (
	New     = "NEW"
	Error   = "ERROR"
	Success = "SUCCESS"
	Failure = "FAILURE"
)

type Payment struct {
	Id        int       `json:"id" example:"35"`
	UserId    int       `json:"user_id" example:"31"`
	UserEmail string    `json:"user_email" example:"user@gmail.com"`
	Amount    int       `json:"amount" example:"3456"`
	Currency  string    `json:"currency" example:"MDA"`
	CreatedAt time.Time `json:"created_at" example:"2022-06-09T14:48:12.288326+03:00"`
	UpdatedAt time.Time `json:"updated_at" example:"2022-06-09T14:48:12.288326+03:00"`
	Status    string    `json:"status" example:"NEW"`
}

func CreatePayment(request NewPaymentRequest) *Payment {
	if request.UserId < 0 {
		return nil
	}
	if request.Amount < 0 {
		return nil
	}
	if _, err := mail.ParseAddress(request.UserEmail); err != nil {
		return nil
	}
	if request.Currency == "" {
		return nil
	}
	return &Payment{
		UserId:    request.UserId,
		UserEmail: request.UserEmail,
		Amount:    request.Amount,
		Currency:  request.Currency,
		Status:    New,
	}
}

func IsStatus(status string) bool {
	return status == New || status == Error || status == Success || status == Failure
}
