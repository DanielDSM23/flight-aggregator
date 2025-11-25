package repos

import (
	"aggregator/domain/models"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Server2Repository struct {
	Listflights []models.Flight
}

func NewServer2Repository() *Server2Repository {
	return &Server2Repository{}
}

func (Server2Repository *Server2Repository) GetFlights() []models.Flight {

	requestURL := "http://localhost:4002/flight_to_book"
	res, err := http.Get(requestURL)
	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
	}
	fmt.Printf("client: got response! Json Server 2\n")
	fmt.Printf("client: status code: Json Server 2 %d\n", res.StatusCode)

	var unMarshalledData []models.Flight
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &unMarshalledData)
	// fmt.Printf("Parsed data Flight2: %#v", unMarshalledData)

	return unMarshalledData
}
