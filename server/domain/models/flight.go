package models

import "time"

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

type Traveler struct {
	FirstName string
	LastName  string
}

type Total struct {
	Amount   float64
	Currency string
}

type Flight struct {
	Reference string
	Status    string
	Traveler  Traveler
	Segments  []Stop
	Total     Total
}

// type Flight struct {
// 	Id        int
// 	Departure string
// 	number    string
// 	from      string
// 	to        string
// 	depart    time.Time
// 	arrive    time.Time
// }

type Flights struct {
	Flights []Flight
}
