package controllers

import (
	"github.com/lyeka/microservices-docker-go-mongodb/movies/models"
)

type (
	// For Get - /movies
	MoviesResource struct {
		Data []models.Movie `json:"data"`
	}
	// For Post/Put - /movies
	MovieResource struct {
		Data models.Movie `json:"data"`
	}
)
