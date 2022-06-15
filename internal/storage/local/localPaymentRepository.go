package local

import (
	"errors"
	"time"

	"github.com/Dshepett/payment-service/internal/models"
)

type PaymentRepository struct {
	Payments []models.Payment
	curr     int
}

func (p *PaymentRepository) AddPayment(payment *models.Payment) error {
	payment.Id = p.curr + 1
	p.curr++
	p.Payments = append(p.Payments, *payment)
	return nil
}

func (p *PaymentRepository) UpdatePaymentStatus(payment *models.Payment) error {
	for pos, val := range p.Payments {
		if val.Id == payment.Id {
			p.Payments[pos].UpdatedAt = time.Now()
			p.Payments[pos].Status = payment.Status
			return nil
		}
	}
	return errors.New("payment with such id does not exist")
}

func (p *PaymentRepository) GetPayment(id int) (*models.Payment, error) {
	for _, val := range p.Payments {
		if val.Id == id {
			return &val, nil
		}
	}
	return nil, errors.New("payment with such id does not exist")
}

func (p *PaymentRepository) GetPaymentsByUserId(id int) ([]models.Payment, error) {
	var payments []models.Payment
	for _, val := range p.Payments {
		if val.UserId == id {
			payments = append(payments, val)
		}
	}
	return payments, nil
}

func (p *PaymentRepository) GetPaymentsByUserEmail(email string) ([]models.Payment, error) {
	var payments []models.Payment
	for _, val := range p.Payments {
		if val.UserEmail == email {
			payments = append(payments, val)
		}
	}
	return payments, nil
}

func (p *PaymentRepository) Close() error {
	return nil
}
