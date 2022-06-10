package app

import (
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/Dshepett/payment-service/pkg/responder"
)

func (a *App) addMiddlewares() {
	a.router.Use(a.logRequest)

}

func (a *App) logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("started %s %s\n", r.Method, r.RequestURI)
		start := time.Now()
		rw := &responseWriter{w, http.StatusOK}
		next.ServeHTTP(rw, r)
		log.Printf("completed with %d %s in %v",
			rw.code,
			http.StatusText(rw.code),
			time.Now().Sub(start))
	})
}

func (a *App) authValidation(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authorizationData := r.Header.Get("Authorization")
		log.Printf("auth: %s\n", authorizationData)
		if authorizationData == "" {
			responder.RespondWithError(w, http.StatusUnauthorized, "authorize before using this function")
			return
		}
		authorizationDataParts := strings.Split(authorizationData, " ")
		if len(authorizationDataParts) != 2 {
			responder.RespondWithError(w, http.StatusUnauthorized, "authorize before using this function")
			return
		}
		username, err := a.service.ParseToken(authorizationDataParts[1])
		if err != nil {
			responder.RespondWithError(w, http.StatusUnauthorized, err.Error())
			return
		}
		log.Printf("User: %s\n", username)
		next.ServeHTTP(w, r)
	}
}
