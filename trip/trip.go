package trip

// Trip
type Trip struct {
	TripID      *int    `json:"id"`
	TimeStamp   string  `json:"time_stamp"`
	UserID      string  `json:"user_id"`
	Distance    int     `json:"distance"`
	GallonsUsed int     `json:"gallons_used"`
	CostOfFuel  float32 `json:"cost_of_fuel"`
	StartDest   string  `json:"start_dest"`
	EndDest     string  `json:"end_dest"`
}
