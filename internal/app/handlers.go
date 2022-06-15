package app

import (
	"encoding/json"
	"net/http"
	"strconv"

	_ "github.com/Dshepett/payment-service/docs"
	"github.com/Dshepett/payment-service/internal/models"
	"github.com/Dshepett/payment-service/pkg/responder"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

func (a *App) addRoutes() {
	a.router.HandleFunc("/auth/login", a.logInHandler).Methods("POST")
	a.router.HandleFunc("/payments/new", a.newPaymentHandler).Methods("POST")
	a.router.HandleFunc("/payments/{id:[0-9]+}/change", a.authValidation(a.changePaymentStatusHandler)).Methods("POST")
	a.router.HandleFunc("/payments/{id:[0-9]+}/status", a.paymentStatusHandler).Methods("GET")
	a.router.HandleFunc("/payments/user/{id:[0-9]+}", a.paymentsByIdHandler).Methods("GET")
	a.router.HandleFunc("/payments/user/email/{email}", a.paymentsByEmailHandler).Methods("GET")
	a.router.HandleFunc("/payments/{id:[0-9]+}/deny", a.denyPaymentHandler).Methods("DELETE")
	a.router.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"),
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("swagger-ui"),
	)).Methods(http.MethodGet)
}

// @Summary SignIn
// @Tags auth
// @Description Use this route to get authorization token that will be needed to change payment's status( Add "Bearer + (token)" in ApiKeyAuth).
// @Accept  json
// @Produce  json
// @Param input body models.LogInRequest true "userdata"
// @Success 200 {object} models.LoginResponse "token"
// @Failure 400 {object} models.ErrorResponse
// @Router /auth/login [post]
func (a *App) logInHandler(w http.ResponseWriter, r *http.Request) {
	var requestData models.LogInRequest
	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		responder.RespondWithError(w, http.StatusBadRequest, "incorrect data")
		return
	}
	token, err := a.service.GenerateToken(requestData.Username, requestData.Password)
	if err != nil {
		responder.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	responder.RespondWithJson(w, http.StatusOK, models.LoginResponse{Token: token})
}

// @Summary      Create new payment
// @Description  Create new payment from request's body.
// @Tags         Payments
// @Accept       json
// @Produce      json
// @Param        input body models.NewPaymentRequest true "payment info"
// @Success      200 {object} models.NewPaymentResponse
// @Failure      400 {object} models.ErrorResponse
// @Router       /payments/new [post]
func (a *App) newPaymentHandler(w http.ResponseWriter, r *http.Request) {
	var requestData models.NewPaymentRequest
	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		responder.RespondWithError(w, http.StatusBadRequest, "incorrect data")
		return
	}
	payment := models.CreatePayment(requestData)
	if payment == nil {
		responder.RespondWithError(w, http.StatusBadRequest, "incorrect data")
		return
	}
	err := a.service.AddPayment(payment)
	if err != nil {
		responder.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	responder.RespondWithJson(w, http.StatusOK, models.NewPaymentResponse{Payment: *payment})
}

// @Summary      Change payment's status
// @Security     ApiKeyAuth
// @Description  Allows to change status if current status is NEW on SUCCESS OR FAILURE. Authorize first!!!!!!!!!!!!!
// @Tags         Payments
// @Accept       json
// @Produce      json
// @Param        id path int true "Payment ID" example(50)
// @Param        input body models.ChangeStatusRequest true "status name"
// @Success      200
// @Failure      400 {object} models.ErrorResponse
// @Router       /payments/{id}/change [post]
func (a *App) changePaymentStatusHandler(w http.ResponseWriter, r *http.Request) {
	var requestData models.ChangeStatusRequest
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		responder.RespondWithError(w, http.StatusBadRequest, "wrong id type")
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		responder.RespondWithError(w, http.StatusBadRequest, "incorrect data")
		return
	}
	if err := a.service.ChangePaymentStatus(id, requestData.Status); err != nil {
		responder.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	responder.RespondWithJson(w, http.StatusOK, nil)
}

// @Summary      Returns payment's status
// @Description  return payment's status if payment exists else returns error message.
// @Tags         Payments
// @Accept       json
// @Produce      json
// @Param        id path int true "Payment ID" example(50)
// @Success      200 {object} models.PaymentStatusResponse
// @Failure      400 {object} models.ErrorResponse
// @Router       /payments/{id}/status [get]
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
	responder.RespondWithJson(w, http.StatusOK, models.PaymentStatusResponse{Status: status})
}

// @Summary      Returns all payments with chosen user ID
// @Description  Returns all payments with chosen user ID.
// @Tags         Payments
// @Accept       json
// @Produce      json
// @Param        id path int true "User ID" example(50)
// @Success      200 {object} models.PaymentsResponse
// @Failure      400 {object} models.ErrorResponse
// @Router       /payments/user/{id} [get]
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
	responder.RespondWithJson(w, http.StatusOK, models.PaymentsResponse{Payments: payments})
}

// @Summary      Returns all payments with chosen user email
// @Description  Returns all payments with chosen user email.
// @Tags         Payments
// @Accept       json
// @Produce      json
// @Param        email path string true "User email" example(user@gmail.com)
// @Success      200 {object} models.PaymentsResponse
// @Failure      400 {object} models.ErrorResponse
// @Router       /payments/user/email/{email} [get]
func (a *App) paymentsByEmailHandler(w http.ResponseWriter, r *http.Request) {
	email := mux.Vars(r)["email"]
	payments, err := a.service.PaymentsByEmail(email)
	if err != nil {
		responder.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	responder.RespondWithJson(w, http.StatusOK, models.PaymentsResponse{Payments: payments})
}

// @Summary      Deny payment if it is possible
// @Description  Deny payment if it exists and its status equals NEW or ERROR and sets DENIED status.
// @Tags         Payments
// @Accept       json
// @Produce      json
// @Param        id path int true "Payment ID" example(50)
// @Success      200
// @Failure      400 {object} models.ErrorResponse
// @Router       /payments/{id}/deny [delete]
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
