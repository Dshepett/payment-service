package app

import (
	"github.com/Dshepett/payment-service/pkg/responder"
	"net/http"
)

func (a *App) addRoutes() {
	a.router.HandleFunc("/", a.helloHandler)
	a.router.HandleFunc("/payments/new", a.newPaymentHandler).Methods("POST")
	a.router.HandleFunc("/payments/{id:[0-9]+}/change", a.changePaymentStatusHandler)
	a.router.HandleFunc("/payments/{id:[0-9]+}/status", a.paymentStatusHandler)
	a.router.HandleFunc("/payments/user/{id:[0-9]+}", a.paymentsByIdHandler)
	a.router.HandleFunc("/payments/user/email/{email}", a.paymentsByEmailHandler)
	a.router.HandleFunc("/payments/{id:[0-9]+}/deny", a.denyPaymentHandler)
}

func (a *App) helloHandler(w http.ResponseWriter, r *http.Request) {
	responder.RespondWithJson(w, http.StatusOK, "hello")
}

func (a *App) newPaymentHandler(w http.ResponseWriter, r *http.Request) {

}

func (a *App) changePaymentStatusHandler(w http.ResponseWriter, r *http.Request) {

}
func (a *App) paymentStatusHandler(w http.ResponseWriter, r *http.Request) {

}

func (a *App) paymentsByIdHandler(w http.ResponseWriter, r *http.Request) {

}

func (a *App) paymentsByEmailHandler(w http.ResponseWriter, r *http.Request) {

}

func (a *App) denyPaymentHandler(w http.ResponseWriter, r *http.Request) {

}
