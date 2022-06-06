package storage

import "github.com/Dshepett/payment-service/internal/config"

type Storage struct {
}

func New(config config.Config) *Storage {
	return &Storage{}
}
