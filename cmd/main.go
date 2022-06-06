package main

import (
	"github.com/Dshepett/payment-service/internal/app"
	"github.com/Dshepett/payment-service/internal/config"
)

func main() {
	conf := config.New()
	application := app.New(conf)
	application.Run()
}
