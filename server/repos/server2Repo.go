package repos

import (
	"aggregator/domain/models"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Server2Repository struct {
	Listflights []models.Flight
}

type StopInfo struct {
	Number string
	From   string
	To     string
	Depart time.Time
	Arrive time.Time
}

type Stop struct {
	Flight StopInfo
}

type TotalServer2 struct {
	Amount   float64
	Currency string
}

type FlightServer2 struct {
	Reference string
	Status    string
	Segments  []Stop
	Total     TotalServer2
}

type Flights struct {
	Flights []FlightServer2
}

func NewServer2Repository() *Server2Repository {
	return &Server2Repository{}
}

func (Server2Repository *Server2Repository) GetFlights() []models.Flight {

	requestURL := "http://j-server2:4002/flight_to_book"
	res, err := http.Get(requestURL)
	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
	}
	fmt.Printf("client: got response! Json Server 2\n")
	fmt.Printf("client: status code: Json Server 2 %d\n", res.StatusCode)
	fmt.Printf("Parsed data Flight2: %#v", res.Body)

	var unMarshalledData []FlightServer2
	body, _ := io.ReadAll(res.Body)

	json.Unmarshal(body, &unMarshalledData)

	var reformatData []models.Flight

	for _, v := range unMarshalledData {
		// fmt.Printf("Parsed data: %#v", v)
		var segments = v.Segments
		stopArray := []models.StopInfo{}

		for _, value := range segments {
			fmt.Printf("segments data: %#v", value)
			var stopInfo = models.StopInfo{Number: value.Flight.Number, From: value.Flight.From, To: value.Flight.To, Depart: value.Flight.Depart, Arrive: value.Flight.Arrive}
			stopArray = append(stopArray, stopInfo)
		}
		total := models.Total{Amount: v.Total.Amount, Currency: v.Total.Currency}

		var flightToAppend = models.Flight{Reference: v.Reference, Status: v.Status, Segments: stopArray, Total: total}
		reformatData = append(reformatData, flightToAppend)
	}

	return reformatData
}
