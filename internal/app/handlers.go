package app

import (
	"net/http"
	"strconv"

	"github.com/Dshepett/payment-service/internal/models"
	"github.com/Dshepett/payment-service/pkg/responder"
	"github.com/gorilla/mux"
)

func (a *App) addRoutes() {
	a.router.HandleFunc("/", a.helloHandler)                                  //
	a.router.HandleFunc("/payments/new", a.newPaymentHandler).Methods("POST") //
	a.router.HandleFunc("/payments/{id:[0-9]+}/change", a.changePaymentStatusHandler).Methods("POST")
	a.router.HandleFunc("/payments/{id:[0-9]+}/status", a.paymentStatusHandler).Methods("GET")
	a.router.HandleFunc("/payments/user/{id:[0-9]+}", a.paymentsByIdHandler).Methods("GET")
	a.router.HandleFunc("/payments/user/email/{email}", a.paymentsByEmailHandler).Methods("GET")
	a.router.HandleFunc("/payments/{id:[0-9]+}/deny", a.denyPaymentHandler).Methods("DELETE")
}

func (a *App) helloHandler(w http.ResponseWriter, r *http.Request) {
	responder.RespondWithJson(w, http.StatusOK, "hello")
}

func (a *App) newPaymentHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		responder.RespondWithError(w, http.StatusBadRequest, "error during reading request's data")
		return
	}
	userId := r.Form.Get("user_id")
	userEmail := r.Form.Get("user_email")
	amount := r.Form.Get("amount")
	currency := r.Form.Get("currency")
	payment := models.CreatePayment(userId, userEmail, amount, currency)
	if payment == nil {
		responder.RespondWithError(w, http.StatusBadRequest, "incorrect data")
		return
	}
	err := a.service.AddPayment(payment)
	if err != nil {
		responder.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	responder.RespondWithJson(w, http.StatusOK, map[string]interface{}{"payment": payment})
}

func (a *App) changePaymentStatusHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		responder.RespondWithError(w, http.StatusBadRequest, "wrong id type")
		return
	}
	if err := r.ParseForm(); err != nil {
		responder.RespondWithError(w, http.StatusBadRequest, "error during reading request's data")
		return
	}
	status := r.Form.Get("status")
	err = a.service.ChangePaymentStatus(id, status)
	if err != nil {
		responder.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	responder.RespondWithJson(w, http.StatusOK, nil)
}

func (a *App) paymentStatusHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		responder.RespondWithError(w, http.StatusBadRequest, "wrong id type")
		return
	}
	status, err := a.service.PaymentStatus(id)
	if err != nil {
		responder.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	responder.RespondWithJson(w, http.StatusOK, map[string]string{"status": status})
}

func (a *App) paymentsByIdHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		responder.RespondWithError(w, http.StatusBadRequest, "wrong id type")
		return
	}
	payments, err := a.service.PaymentsByUserId(id)
	if err != nil {
		responder.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	responder.RespondWithJson(w, http.StatusOK, map[string]interface{}{"payments": payments})
}

func (a *App) paymentsByEmailHandler(w http.ResponseWriter, r *http.Request) {
	email := mux.Vars(r)["email"]
	payments, err := a.service.PaymentsByEmail(email)
	if err != nil {
		responder.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	responder.RespondWithJson(w, http.StatusOK, map[string]interface{}{"payments": payments})
}

func (a *App) denyPaymentHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		responder.RespondWithError(w, http.StatusBadRequest, "wrong id type")
		return
	}
	err = a.service.DenyPayment(id)
	if err != nil {
		responder.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	responder.RespondWithJson(w, http.StatusOK, nil)
}
