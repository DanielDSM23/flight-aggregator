package domain

import (
	"aggregator/application"
	"aggregator/domain/models"
)

type Service struct {
	handler *application.Handler
}

func NewService(handler *application.Handler) *Service {
	return &Service{}
}

func (s *Service) SortByPrice() models.Flights {

	return models.Flights{}
}

func (s *Service) SortByTimeTravel() models.Flights {
	return models.Flights{}
}

func (s *Service) SortByDepartureDate() models.Flights {
	return models.Flights{}
}
