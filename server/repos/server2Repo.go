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
	Listflights []FlightServer2
}

type StopInfo struct {
	Number string    `json:"number"`
	From   string    `json:"from"`
	To     string    `json:"to"`
	Depart time.Time `json:"depart"`
	Arrive time.Time `json:"arrive"`
}

type Stop struct {
	Flight StopInfo `json:"flight"`
}

type Traveler struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type Total struct {
	Amount   float32 `json:"amount"`
	Currency string  `json:"currency"`
}

type FlightServer2 struct {
	Reference string   `json:"reference"`
	Status    string   `json:"status"`
	Traveler  Traveler `json:"traveler"`
	Segments  []Stop   `json:"segments"`
	Total     Total    `json:"total"`
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

	var unMarshalledData []FlightServer2
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &unMarshalledData)

	fmt.Printf("Parsed data: %#v", unMarshalledData)
	var dataFormatted = models.Flights{}

	return dataFormatted
}
