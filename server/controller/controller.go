package controller

import (
	"aggregator/application"
	"aggregator/repos"
	"encoding/json"
	"fmt"
	"net/http"
)

type Controller struct {
	// server *http.ServeMux
}

func NewController() *Controller {
	return &Controller{}
}

func (controller *Controller) GetflightsServer1(responseWriter http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(responseWriter, "Hello there!")
	FlightServer1Repo := repos.NewServer1Repository()
	// FlightServer2Repo := repos.NewServer2Repository()
	handler := application.NewHandler(FlightServer1Repo)

	data := handler.CombineData()
	fmt.Fprintf(responseWriter, "data =>", data)
	dataJson, _ := json.MarshalIndent(data, "", "  ")
	fmt.Println(string(dataJson))
}
