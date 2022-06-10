package local

import (
	"github.com/Dshepett/payment-service/internal/storage"
)

type Storage struct {
	PaymentRepository storage.PaymentRepository
}

func (s *Storage) Open() error {
	s.PaymentRepository = &PaymentRepository{}
	return nil
}

func (s *Storage) Payment() storage.PaymentRepository {
	return s.PaymentRepository
}

func (s *Storage) Close() error {
	return nil
}
