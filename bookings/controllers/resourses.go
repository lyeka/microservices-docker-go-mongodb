package controllers

import (
	"github.com/lyeka/microservices-docker-go-mongodb/bookings/models"
)

type (
	// For Get - /bookings
	BookingsResource struct {
		Data []models.Booking `json:"data"`
	}
	// For Post/Put - /bookings
	BookingResource struct {
		Data models.Booking `json:"data"`
	}
)
