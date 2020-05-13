package main

import (
	"log"
	"net/http"

	"github.com/lyeka/microservices-docker-go-mongodb/users/common"
	"github.com/lyeka/microservices-docker-go-mongodb/users/routers"
)

// Entry point for the program
func main() {

	// Calls startup logic
	common.StartUp()
	// Get the mux router object
	router := routers.InitRoutes()

	server := &http.Server{
		Addr:    common.AppConfig.Server,
		Handler: router,
	}
	log.Println("Listening...")
	server.ListenAndServe()
}
