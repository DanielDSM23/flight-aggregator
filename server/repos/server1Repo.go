package repos

import (
	"aggregator/domain/models"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

type Server1Repository struct {
	Listflights []FlightServer1
}

type FlightServer1 struct {
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

func NewServer1Repository() *Server1Repository {
	return &Server1Repository{}
}

func (Server1Repository *Server1Repository) GetFlights() []models.Flight {

	requestURL := "http://localhost:4001/flights"
	res, err := http.Get(requestURL)
	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
	}
	fmt.Printf("client: got response! Json Server 1\n")
	fmt.Printf("client: status code: Json Server 1 %d\n", res.StatusCode)

	var unMarshalledData []FlightServer1
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &unMarshalledData)

	var dataFormatted = []models.Flight{}

	for _, v := range unMarshalledData {
		stringSlice := strings.Split(v.PassengerName, " ")
		var firstname = stringSlice[0]
		var lastname string

		for j, value := range stringSlice {
			if j > 0 && len(stringSlice) > 2 {
				lastname += value + " "
			} else if j > 0 && len(stringSlice) == 2 {
				lastname += value
			}
		}

		var stopInfo = models.StopInfo{Number: v.FlightNumber, From: v.DepartureAirport, To: v.ArrivalAirport, Depart: v.DepartureTime, Arrive: v.ArrivalTime}

		var stop = models.Stop{Flight: stopInfo}
		stopArray := []models.Stop{}
		stopArray = append(stopArray, stop)

		var travelers = models.Traveler{FirstName: firstname, LastName: lastname}
		var total = models.Total{Amount: v.Price, Currency: v.Currency}

		var flightToAppend = models.Flight{Reference: v.BookingId, Status: v.Status, Traveler: travelers, Segments: stopArray, Total: total}
		dataFormatted = append(dataFormatted, flightToAppend)
	}

	// fmt.Printf("Parsed data: %#v", dataFormatted)
	return dataFormatted
}
