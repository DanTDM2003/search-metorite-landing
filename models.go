package main

type GeoLocation struct {
	Latitude      string `json:"latitude"`
	Longitude     string `json:"longitude"`
	NeedsRecoding bool   `json:"needs_recoding"`
}

type MetoriteLanding struct {
	ID          string      `json:"id"`
	Year        string      `json:"year"`
	Name        string      `json:"name"`
	Nametype    string      `json:"nametype"`
	Recclass    string      `json:"recclass"`
	Mass        string      `json:"mass"`
	Fall        string      `json:"fall"`
	Reclat      string      `json:"reclat"`
	Reclong     string      `json:"reclong"`
	GeoLocation GeoLocation `json:"geolocation"`
}
