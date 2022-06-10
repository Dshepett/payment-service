package storage

type Storage interface {
	Open() error
	Payment() PaymentRepository
	Close() error
}
