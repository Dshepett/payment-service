package service

import (
	"errors"
	"log"
	"math/rand"
	"time"

	"github.com/Dshepett/payment-service/internal/config"
	"github.com/Dshepett/payment-service/internal/models"
	"github.com/Dshepett/payment-service/internal/storage"
)

type Service struct {
	storage *storage.Storage
}

func New(config *config.Config) *Service {
	return &Service{storage: storage.New(config)}
}

func (s *Service) Start() {
	err := s.storage.Open()
	if err != nil {
		log.Fatal(err)
	}
}

func (s *Service) AddPayment(payment *models.Payment) error {
	payment.CreatedAt = time.Now()
	payment.UpdatedAt = time.Now()
	if rand.ExpFloat64() > 0.5 {
		payment.Status = models.Error
	}
	err := s.storage.Payment().AddPayment(payment)
	return err
}

func (s *Service) ChangePaymentStatus(id int, status string) error {
	if !models.IsStatus(status) {
		return errors.New("such status does not exist")
	}
	payment, err := s.storage.Payment().GetPayment(id)
	if err != nil {
		return errors.New("payment with such id does not exist")
	}
	if payment.Status != models.New {
		return errors.New("status could not be changed")
	}
	payment.Status = status
	payment.UpdatedAt = time.Now()
	err = s.storage.Payment().UpdatePaymentStatus(payment)
	if err != nil {
		return errors.New("error occurred during changing status")
	}
	return nil
}

func (s *Service) PaymentStatus(id int) (string, error) {
	if payment, err := s.storage.Payment().GetPayment(id); err != nil {
		return "", errors.New("payment with such id does not exist")
	} else {
		return payment.Status, nil
	}
}

func (s *Service) PaymentsByUserId(id int) ([]models.Payment, error) {
	payments, err := s.storage.Payment().GetPaymentsByUserId(id)
	if err != nil {
		return nil, errors.New("error occurred during finding payments")
	}
	return payments, nil
}

func (s *Service) PaymentsByEmail(email string) ([]models.Payment, error) {
	payments, err := s.storage.Payment().GetPaymentsByUserEmail(email)
	if err != nil {
		return nil, errors.New("error occurred during finding payments")
	}
	return payments, nil
}

func (s *Service) DenyPayment(id int) error {
	payment, err := s.storage.Payment().GetPayment(id)
	if err != nil {
		return errors.New("payment with such id does not exist")
	}
	if payment.Status == models.Success || payment.Status == models.Failure {
		return errors.New("this payment could not be denied")
	}
	err = s.storage.Payment().DeletePaymentById(id)
	if err != nil {
		return errors.New("error occurred during denying this payment")
	}
	return nil
}

func (s *Service) Close() error {
	return s.storage.Close()
}
