package routers

import (
	"github.com/gorilla/mux"
	"github.com/lyeka/microservices-docker-go-mongodb/bookings/controllers"
)

func SetBookingsRouters(router *mux.Router) *mux.Router {
	router.HandleFunc("/bookings", controllers.GetBookings).Methods("GET")
	router.HandleFunc("/bookings", controllers.CreateBooking).Methods("POST")
	return router
}
