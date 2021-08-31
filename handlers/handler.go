package handlers

import (
	"net/http"

	"github.com/pmadhvi/iban-validator/validator"
)

func validateIban(rw http.ResponseWriter, req *http.Request) {
	// read iban from request url
	iban := req.URL.Query().Get("iban")

	err := validator.IbanValidator(iban)
	if err != nil {
		respondSuccessJSON(rw, []byte(`{"message": "iban is invalid"}`))
	}
	respondSuccessJSON(rw, []byte(`{"message": "iban is valid"}`))
}

func checkHealth(rw http.ResponseWriter, req *http.Request) {
	respondSuccessJSON(rw, []byte(`{"message": "alive"}`))
}

func InternalServerwrror(rw http.ResponseWriter, req *http.Request) {
	errMsg := []byte(`{"error_message": "server is down"}`)
	respondErrorJSON(rw, http.StatusInternalServerError, errMsg)
}

func respondSuccessJSON(rw http.ResponseWriter, response []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))
}

func respondErrorJSON(rw http.ResponseWriter, errorCode int, errorMsg []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(errorCode)
	w.Write([]byte(errorMsg))
}
