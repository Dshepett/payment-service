package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/Dshepett/payment-service/internal/app"
	"github.com/Dshepett/payment-service/internal/config"
)

// @title           Payment Service APi
// @version         1.0
// @description   	Simple api for handling payments.

// @host      localhost:8080
// @BasePath  /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	conf := config.New()
	application := app.New(conf)
	go application.Run()
	exit := make(chan os.Signal, 1)
	signal.Notify(exit, syscall.SIGINT, syscall.SIGTERM)
	<-exit
	application.Close()
}
