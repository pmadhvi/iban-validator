package main

import (
	"fmt"
	"net/http"
	"os"
	"syscall"

	"github.com/pmadhvi/iban-validator/handlers"
	"github.com/sirupsen/logrus"
)

func main() {
	// setup the log
	var log = logrus.New()
	log.SetOutput(os.Stdout)

	// read the port from env variable
	port, ok := syscall.Getenv("port")
	if !ok {
		log.Info("port env variable not set, so using default port 8080")
		port = "8080"
	}

	// setup router
	router := handlers.Router{
		Log: log,
	}

	// setup all the routes
	http.Handle("/", router.Routes())

	// start the server on specified port
	log.Fatalf("server failed to start: %v", http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
