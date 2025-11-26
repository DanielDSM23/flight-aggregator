package domain

import (
	"aggregator/domain/models"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRepo struct {
	mock.Mock
}

func (m *MockRepo) GetFlights() []models.Flight {
	args := m.Called()
	return args.Get(0).([]models.Flight)
}

func createFlight(total float64, departure time.Time, arrival time.Time) models.Flight {
	return models.Flight{
		Total: models.Total{Amount: total},
		Segments: []models.Stop{
			{
				Flight: models.StopInfo{
					Depart: departure,
					Arrive: arrival,
				},
			},
		},
	}
}

func TestSortByPrice(t *testing.T) {
	service := NewService(nil)
	flights := []models.Flight{
		createFlight(200.0, time.Now(), time.Now()),
		createFlight(100.0, time.Now(), time.Now()),
		createFlight(150.0, time.Now(), time.Now()),
	}

	service.SortByPrice(&flights)

	assert.Equal(t, 100.0, flights[0].Total.Amount, "100 expected")
	assert.Equal(t, 150.0, flights[1].Total.Amount, "150 expected")
	assert.Equal(t, 200.0, flights[2].Total.Amount, "200 expected")
}
