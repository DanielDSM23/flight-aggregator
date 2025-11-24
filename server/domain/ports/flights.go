package ports

import "aggregator/domain/models"

type Repos interface {
	GetFlights() models.Flights
}
