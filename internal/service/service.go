package service

import (
	"github.com/Dshepett/payment-service/internal/config"
	"github.com/Dshepett/payment-service/internal/models"
	"github.com/Dshepett/payment-service/internal/storage"
)

type Service struct {
	storage *storage.Storage
}

func New(config config.Config) *Service {
	return &Service{storage: nil}
}

func (s *Service) AddPayment(payment models.Payment) error {
	return nil
}

func changePaymentStatus(id int, status string) error {
	return nil
}

func PaymentStatus(id int) (string, error) {
	return "", nil
}

func PaymentsByUserId(id int) ([]models.Payment, error) {
	return nil, nil
}

func PaymentsByEmail(email string) ([]models.Payment, error) {
	return nil, nil
}

func denyPayment(id int) error {
	return nil
}
