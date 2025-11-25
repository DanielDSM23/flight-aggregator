package routing

import (
	"aggregator/controller"
	"net/http"
)

func Router() *http.ServeMux {

	mux := http.NewServeMux()
	controller := controller.NewController()
	mux.HandleFunc("GET /flight", controller.GetCombinedflights)
	return mux
}
