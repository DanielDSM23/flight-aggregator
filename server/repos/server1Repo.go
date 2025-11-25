package repos

import (
	"aggregator/domain/models"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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

func (Server1Repository *Server1Repository) GetFlights() models.Flights {

	requestURL := "http://localhost:4001/flights"
	fmt.Printf("test \n")

	res, err := http.Get(requestURL)
	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
	}
	fmt.Printf("client: got response!\n")
	fmt.Printf("client: status code: %d\n", res.StatusCode)

	var unMarshalledData []FlightServer1
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &unMarshalledData)

	fmt.Printf("request result: %#v", unMarshalledData)

	var dataFormatted = models.Flights{}
	// for
	// errorRequest := json.NewDecoder(res.Body).Decode(&dataFormatted)
	// if errorRequest != nil {
	// 	fmt.Println(errorRequest)
	// 	return dataFormatted
	// }

	// fmt.Printf("Parsed data: %+v", dataFormatted)
	return dataFormatted
}
