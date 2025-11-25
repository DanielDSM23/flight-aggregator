package controller

import (
	"aggregator/application"
	"aggregator/domain/ports"
	domain "aggregator/domain/services"
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

func (controller *Controller) GetCombinedflights(responseWriter http.ResponseWriter, request *http.Request) {

	FlightServer1Repo := repos.NewServer1Repository()
	FlightServer2Repo := repos.NewServer2Repository()
	repos := []ports.Repos{}
	repos = append(repos, FlightServer1Repo)
	repos = append(repos, FlightServer2Repo)

	handler := application.NewHandler(repos)

	data := handler.CombineData()
	domain := domain.NewService(nil)
	filterValue := request.URL.Query().Get("filter")
	switch filterValue {
	case "SortByPrice":
		domain.SortByPrice(&data)
	case "SortByTimeTravel":
		domain.SortByTimeTravel(&data)
	case "SortByDepartureDate":
		domain.SortByDepartureDate(&data)
	}
	var val, _ = json.Marshal(data)
	var jsonData = string(val)
	fmt.Fprintf(responseWriter, "%s", jsonData)
}
