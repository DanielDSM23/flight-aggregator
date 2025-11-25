package controller

import (
	"aggregator/application"
	"aggregator/domain/ports"
	"aggregator/repos"
	"fmt"
	"net/http"
)

type Controller struct {
	// server *http.ServeMux
}

func NewController() *Controller {
	return &Controller{}
}

func (controller *Controller) GetCombinedflights(responseWriter http.ResponseWriter, request *http.Request) {
	// fmt.Fprintf(responseWriter, "Hello there!")
	FlightServer1Repo := repos.NewServer1Repository()
	FlightServer2Repo := repos.NewServer2Repository()
	repos := []ports.Repos{}
	repos = append(repos, FlightServer1Repo)
	repos = append(repos, FlightServer2Repo)

	handler := application.NewHandler(repos)

	data := handler.CombineData()
	fmt.Println(responseWriter, "data =>", data)
	// dataJson, _ := json.MarshalIndent(data, "", "  ")
	// fmt.Println(string(dataJson))
}
