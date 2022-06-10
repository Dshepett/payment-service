package responder

import (
	"encoding/json"
	"net/http"

	"github.com/Dshepett/payment-service/internal/models"
)

func RespondWithJson(w http.ResponseWriter, code int, data interface{}) {
	response, _ := json.Marshal(data)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func RespondWithError(w http.ResponseWriter, code int, message string) {
	RespondWithJson(w, code, models.ErrorResponse{Message: message})
}
