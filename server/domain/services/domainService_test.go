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
		Segments: []models.StopInfo{
			{

				Depart: departure,
				Arrive: arrival,
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

func TestSortByTimeTravel(t *testing.T) {
	service := NewService(nil)
	now := time.Now()
	flights := []models.Flight{
		createFlight(100, now, now.Add(3*time.Hour)),
		createFlight(100, now, now.Add(1*time.Hour)),
		createFlight(100, now, now.Add(2*time.Hour)),
	}
	service.SortByTimeTravel(&flights)

	assert.Equal(t, 1*time.Hour, flights[0].Segments[0].Arrive.Sub(flights[0].Segments[0].Depart), "Should be equal to 1")
	assert.Equal(t, 2*time.Hour, flights[1].Segments[0].Arrive.Sub(flights[1].Segments[0].Depart), "Should be equal to 2")
	assert.Equal(t, 3*time.Hour, flights[2].Segments[0].Arrive.Sub(flights[2].Segments[0].Depart), "Should be equald to 3")
}

func TestSortByDepartureDate(t *testing.T) {
	service := NewService(nil)
	now := time.Now()
	flights := []models.Flight{
		createFlight(100, now.Add(3*time.Hour), now.Add(4*time.Hour)),
		createFlight(100, now.Add(1*time.Hour), now.Add(2*time.Hour)),
		createFlight(100, now.Add(2*time.Hour), now.Add(3*time.Hour)),
	}

	service.SortByDepartureDate(&flights)

	assert.True(t, flights[0].Segments[0].Depart.Before(flights[1].Segments[0].Depart), "Should be true")
	assert.True(t, flights[1].Segments[0].Depart.Before(flights[2].Segments[0].Depart), "Should be true")
	assert.Equal(t, now.Add(1*time.Hour), flights[0].Segments[0].Depart)
}
