package models

import (
	"net/mail"
	"strconv"
	"time"
)

var (
	New     = "NEW"
	Error   = "ERROR"
	Success = "SUCCESS"
	Failure = "FAILURE"
)

type Payment struct {
	Id        int       `json:"id"`
	UserId    int       `json:"user_id"`
	UserEmail string    `json:"user_email"`
	Amount    int       `json:"amount"`
	Currency  string    `json:"currency"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Status    string    `json:"status"`
}

func CreatePayment(userId string, userEmail string, amount string, currency string) *Payment {
	id, err := strconv.Atoi(userId)
	if err != nil || id < 0 {
		return nil
	}
	amountInt, err := strconv.Atoi(amount)
	if err != nil || amountInt < 0 {
		return nil
	}
	if _, err = mail.ParseAddress(userEmail); err != nil {
		return nil
	}
	return &Payment{
		UserId:    id,
		UserEmail: userEmail,
		Amount:    amountInt,
		Currency:  currency,
		Status:    New,
	}
}

func IsStatus(status string) bool {
	return status == New || status == Error || status == Success || status == Failure
}
