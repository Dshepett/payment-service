package app

import (
	"github.com/Dshepett/payment-service/internal/config"
	"github.com/Dshepett/payment-service/internal/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type App struct {
	service *service.Service
	router  *mux.Router
}

func New(config *config.Config) *App {
	app := &App{
		service: nil,
		router:  mux.NewRouter(),
	}
	app.addRoutes()
	return app
}

func (a *App) Run() {
	s := &http.Server{
		Addr:         "127.0.0.1:8080",
		Handler:      a.router,
		ReadTimeout:  0,
		WriteTimeout: 0,
	}
	log.Fatal(s.ListenAndServe())
}
