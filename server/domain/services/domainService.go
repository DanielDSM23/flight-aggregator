package domain

import (
	"aggregator/application"
	"aggregator/domain/models"
	"sort"
)

type Service struct {
	handler *application.Handler
}

func NewService(handler *application.Handler) *Service {
	return &Service{}
}

func (s *Service) SortByPrice(data *[]models.Flight) {

	sort.Slice(*data, func(i, j int) bool {
		return (*data)[i].Total.Amount < (*data)[j].Total.Amount
	})

}

func (s *Service) SortByTimeTravel(data *[]models.Flight) {
	sort.Slice(*data, func(i, j int) bool {
		flightA := (*data)[i]
		flightB := (*data)[j]
		lastSegmentIndexA := len(flightA.Segments) - 1
		durationA := flightA.Segments[lastSegmentIndexA].Flight.Arrive.Sub(
			flightA.Segments[0].Flight.Depart,
		)
		lastSegmentIndexB := len(flightB.Segments) - 1
		durationB := flightB.Segments[lastSegmentIndexB].Flight.Arrive.Sub(
			flightB.Segments[0].Flight.Depart,
		)
		return durationA < durationB
	})

}

func (s *Service) SortByDepartureDate(data *[]models.Flight) {
	sort.Slice(*data, func(i, j int) bool {
		departTimeI := (*data)[i].Segments[0].Flight.Depart
		departTimeJ := (*data)[j].Segments[0].Flight.Depart

		return departTimeI.Before(departTimeJ)
	})
}
