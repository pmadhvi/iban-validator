package handlers

import (
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

type Router struct {
	Log *log.Logger
}

//Router define handlers based on path
func (r Router) Routes() *mux.Router {
	// Initialize mux router
	router := mux.NewRouter()

	// Define routes and call their handler function
	router.HandleFunc("/api/iban/health", checkHealth)
	router.HandleFunc("/api/iban/validate/{iban}", validateIban)

	return router
}
