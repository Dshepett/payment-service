package storage

import (
	"database/sql"
	"fmt"

	"github.com/Dshepett/payment-service/internal/config"
	"github.com/Dshepett/payment-service/internal/storage/postgres"
)

type Storage struct {
	config            *config.Config
	paymentRepository PaymentRepository
}

func New(config *config.Config) *Storage {
	return &Storage{config: config}
}

func (s *Storage) Open() error {
	connectionString := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		s.config.DBUser, s.config.DBPassword, s.config.DBName)
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return err
	}
	s.paymentRepository = postgres.NewPaymentRepository(db)
	return nil
}

func (s *Storage) Payment() PaymentRepository {
	return s.paymentRepository
}
