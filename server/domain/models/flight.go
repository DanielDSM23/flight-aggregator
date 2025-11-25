package models

import "time"

type StopInfo struct {
	number string
	from   string
	to     string
	depart time.Time
	arrive time.Time
}

type Stop struct {
	flight StopInfo
}

type Traveler struct {
	firstName string
	lastName  string
}

type Total struct {
	amount   float32
	currency string
}

type Flight struct {
	reference string
	status    string
	traveler  Traveler
	segments  []Stop
	total     Total
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
