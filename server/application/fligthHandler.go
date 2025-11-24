package application

import (
	"aggregator/domain/models"
	"aggregator/domain/ports"
)

type Handler struct {
	repos []ports.Repos
}

func NewHandler(repos ...ports.Repos) *Handler {
	return &Handler{
		repos: repos,
	}
}

func (handler *Handler) CombineData() []models.Flight {
	allFlights := []models.Flight{}
	repos := handler.repos
	for _, r := range repos {
		flights := r.GetFlights()

		allFlights = append(allFlights, flights.Flights...)
	}
	return allFlights
}
