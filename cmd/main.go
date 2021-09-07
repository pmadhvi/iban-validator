package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/pmadhvi/iban-validator/handlers"
	"github.com/sirupsen/logrus"
)

func main() {
	// setup the log
	var log = logrus.New()
	log.SetOutput(os.Stdout)

	// read .env file for env variables
	err := godotenv.Load()
	if err != nil {
		log.Error("Error loading .env file", err)
	}
	// read the port from env variable
	port, ok := syscall.Getenv("PORT")
	if !ok {
		log.Info("port env variable not set, so using default port 8080")
		port = "8080"
	}

	// setup server and routes
	server := handlers.Server{
		Log:  log,
		Port: port,
	}

	errorChan := make(chan error)

	// setup all the routes and start the server
	go func() {
		errorChan <- server.Start()
	}()

	// catch the exit signals and pass on errorChan
	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		errorChan <- fmt.Errorf("Got quit signal: %s", <-quit)
	}()

	// get errors from chan and exit the application
	if err := <-errorChan; err != nil {
		log.Error(err)
		os.Exit(1)
	}

}
