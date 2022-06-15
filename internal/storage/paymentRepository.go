package storage

import "github.com/Dshepett/payment-service/internal/models"

type PaymentRepository interface {
	AddPayment(payment *models.Payment) error
	UpdatePaymentStatus(payment *models.Payment) error
	GetPayment(id int) (*models.Payment, error)
	GetPaymentsByUserId(id int) ([]models.Payment, error)
	GetPaymentsByUserEmail(email string) ([]models.Payment, error)
	Close() error
}
