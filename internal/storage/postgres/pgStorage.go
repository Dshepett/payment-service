package postgres

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Dshepett/payment-service/internal/config"
	"github.com/Dshepett/payment-service/internal/storage"
)

type PgStorage struct {
	config            *config.Config
	paymentRepository storage.PaymentRepository
}

func NewStorage(config *config.Config) *PgStorage {
	return &PgStorage{config: config}
}

func (s *PgStorage) Open() error {
	log.Println("connecting to database...")
	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		"postgres", 5432, s.config.DBUser, s.config.DBPassword, s.config.DBName)
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return err
	}
	s.paymentRepository = NewPaymentRepository(db)
	log.Println("success connecting to database")
	return nil
}

func (s *PgStorage) Payment() storage.PaymentRepository {
	return s.paymentRepository
}

func (s *PgStorage) Close() error {
	return s.Payment().Close()
}
