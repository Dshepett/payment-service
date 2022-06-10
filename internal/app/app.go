package app

import (
	"context"
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
	server  *http.Server
}

func New(config *config.Config) *App {
	app := &App{
		service: service.New(config),
		router:  mux.NewRouter(),
	}
	app.addRoutes()
	log.Println("app created")
	return app
}

func (a *App) Run() {
	log.Println("starting server...")
	s := &http.Server{
		Addr:         ":8080",
		Handler:      a.router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}
	a.server = s
	a.service.Start()

	log.Fatal(a.server.ListenAndServe())
}

func (a *App) Close() {
	log.Println("closing ...")
	if err := a.service.Close(); err != nil {
		log.Println("error occurred during shutting down server")
	}

	if err := a.server.Shutdown(context.Background()); err != nil {
		log.Println("error occurred during shutting down database connection")
	}
	log.Println("app closed")
}
