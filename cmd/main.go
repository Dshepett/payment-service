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
// @description     This is a sample server celler server.

// @host      localhost:8080
// @BasePath  /

func main() {
	conf := config.New()
	application := app.New(conf)
	go application.Run()
	exit := make(chan os.Signal, 1)
	signal.Notify(exit, syscall.SIGINT, syscall.SIGTERM)
	<-exit
	application.Close()
}
