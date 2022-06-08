package app

import (
	"log"
	"net/http"
	"time"

	"github.com/Dshepett/payment-service/internal/config"
	"github.com/Dshepett/payment-service/internal/service"
	"github.com/gorilla/mux"
)

type App struct {
	service *service.Service
	router  *mux.Router
	config  *config.Config
}

func New(config *config.Config) *App {
	app := &App{
		service: service.New(config),
		router:  mux.NewRouter(),
		config:  config,
	}
	app.addRoutes()
	return app
}

func (a *App) Run() {
	s := &http.Server{
		Addr:         ":" + a.config.Port,
		Handler:      a.router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}
	a.service.Start()
	log.Fatal(s.ListenAndServe())
}
