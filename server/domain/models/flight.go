package models

import "time"

type Flight struct {
	Id        int
	Departure string
	number    string
	from      string
	to        string
	depart    time.Time
	arrive    time.Time
}

type Flights struct {
	Flights []Flight
}
