package trip

// Trip
type Trip struct {
	TripID      *int   `json:"id"`
	TimeStamp   string `json:"time_stamp"`
	Name        string `json:"name"`
	Distance    int    `json:"distance"`
	GallonsUsed int    `json:"gallons_used"`
	CostOfFuel  int    `json:"cost_of_fuel"`
	Flight      string `json:"flight"`
}

// Data displayed
type DataDisplay struct {
	Name              string  `json:"name"`
	DistanceTraveled  string  `json:"distance_traveled"`
	CO2Produced       float64 `json:"co2_produced"`
	PlasticStrawsUsed float64 `json:"plastic_straws_used"`
}
