package application

import (
	"aggregator/domain/models"
	"aggregator/domain/ports"
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

func createFlight(reference string, total float64, departure time.Time, arrival time.Time) models.Flight {
	return models.Flight{
		Reference: reference,
		Total:     models.Total{Amount: total},
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

func TestCombineData(t *testing.T) {
	mockRepo1 := new(MockRepo)
	mockRepo2 := new(MockRepo)

	flightsFromRepo1 := []models.Flight{
		createFlight("1", 100.0, time.Now().Add(time.Hour), time.Now().Add(2*time.Hour)),
		createFlight("2", 200.0, time.Now().Add(2*time.Hour), time.Now().Add(3*time.Hour)),
	}
	flightsFromRepo2 := []models.Flight{
		createFlight("3", 150.0, time.Now().Add(3*time.Hour), time.Now().Add(4*time.Hour)),
		createFlight("4", 250.0, time.Now().Add(4*time.Hour), time.Now().Add(5*time.Hour)),
	}

	mockRepo1.On("GetFlights").Return(flightsFromRepo1)
	mockRepo2.On("GetFlights").Return(flightsFromRepo2)

	handler := NewHandler([]ports.Repos{mockRepo1, mockRepo2})

	combinedFlights := handler.CombineData()

	mockRepo1.AssertExpectations(t)
	mockRepo2.AssertExpectations(t)

	assert.Len(t, combinedFlights, 4, "Should contain 4 flight")

	for _, f := range flightsFromRepo1 {
		assert.Contains(t, combinedFlights, f, "Should contain repo 1 flight")
	}
	for _, f := range flightsFromRepo2 {
		assert.Contains(t, combinedFlights, f, "Should contain repo 2 flight")
	}
}
