package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pmadhvi/iban-validator/validator"
	log "github.com/sirupsen/logrus"
)

// Response for returning the response of validation test
type Response struct {
	Message string `json:"message"`
}

// ValidateIbanHandler is an httphandler to handle request to validate iban
func ValidateIbanHandler(rw http.ResponseWriter, req *http.Request) {

	// feteching the quary parameters from request url
	vars := mux.Vars(req)
	iban := vars["iban"]

	err := validator.IbanValidation(iban)
	if err != nil {
		errMsg := Response{
			Message: err.Error(),
		}
		log.Errorf("Error response: %v", errMsg)
		respondErrorJSON(rw, 200, errMsg)
		return
	}
	respMsg := Response{
		Message: "Iban is valid.",
	}
	log.Infof("Success response: %v", respMsg)
	respondSuccessJSON(rw, respMsg)
}

// CheckHealthHandler is an httphandler to handle request to check application health
func CheckHealthHandler(rw http.ResponseWriter, req *http.Request) {
	respMsg := Response{
		Message: "Iban Validator application is alive.",
	}
	log.Infof("Health check response: %v", respMsg)
	respondSuccessJSON(rw, respMsg)
}

func respondSuccessJSON(rw http.ResponseWriter, response interface{}) {
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(response)
}

func respondErrorJSON(rw http.ResponseWriter, errorCode int, errorMsg interface{}) {
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(errorCode)
	json.NewEncoder(rw).Encode(errorMsg)
}
