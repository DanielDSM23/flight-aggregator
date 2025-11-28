package models

import "time"

type StopInfo struct {
	Number string
	From   string
	To     string
	Depart time.Time
	Arrive time.Time
}

type Total struct {
	Amount   float64
	Currency string
}

type Flight struct {
	Reference string
	Status    string
	Segments  []StopInfo
	Total     Total
}

type Flights struct {
	Flights []Flight
}
