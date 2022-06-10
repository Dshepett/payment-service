package storage

import (
	"database/sql"
	"fmt"
	"log"

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
	log.Println("connecting to database...")
	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		"postgres", 5432, s.config.DBUser, s.config.DBPassword, s.config.DBName)
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return err
	}
	s.paymentRepository = postgres.NewPaymentRepository(db)
	log.Println("success connecting to database")
	return nil
}

func (s *Storage) Payment() PaymentRepository {
	return s.paymentRepository
}

func (s *Storage) Close() error {
	return s.Payment().Close()
}
