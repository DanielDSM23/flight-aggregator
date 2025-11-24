package repos

import (
	"aggregator/domain/models"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Server2Repository struct {
	Listflights []FlightServer1
}

type FlightServer2 struct {
	Id               int       `json:"id"`
	BookingId        string    `json:"bookingId"`
	Status           string    `json:"status"`
	PassengerName    string    `json:"passengerName"`
	FlightNumber     string    `json:"flightNumber"`
	DepartureAirport string    `json:"departureAirport"`
	ArrivalAirport   string    `json:"arrivalAirport"`
	DepartureTime    time.Time `json:"departureTime"`
	ArrivalTime      time.Time `json:"arrivalTime"`
	Price            float64   `json:"price"`
	Currency         string    `json:"currency"`
}

func NewServer2Repository() *Server2Repository {
	return &Server2Repository{}
}

func (Server2Repository *Server2Repository) GetFlights() models.Flights {

	requestURL := "http://j-server1:4001/flights"
	res, err := http.Get(requestURL)
	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
	}
	fmt.Printf("client: got response!\n")
	fmt.Printf("client: status code: %d\n", res.StatusCode)

	var dataFormatted = models.Flights{}
	errorRequest := json.NewDecoder(res.Body).Decode(&dataFormatted)
	if errorRequest != nil {
		fmt.Println(errorRequest)
		return dataFormatted
	}
	fmt.Printf("Parsed data: %+v", dataFormatted)
	return dataFormatted
}
