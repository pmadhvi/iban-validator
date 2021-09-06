package handlers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

type Server struct {
	Log  *log.Logger
	Port string
}

//  Defines routes and their handlers and start the server
func (s Server) Start() error {
	// Initialize mux router
	router := mux.NewRouter()
	// Define routes and call their handler function
	router.HandleFunc("/api/iban/validate/health", CheckHealthHandler)
	router.HandleFunc("/api/iban/validate/{iban}", ValidateIbanHandler)

	// start the server on specified port
	err := http.ListenAndServe(fmt.Sprintf(":%s", s.Port), router)
	log.Errorf("error starting server: %v", err)
	return err
}
