package trip

// Trip
type Trip struct {
	TripID      *int
	TimeStamp   string
	UserID      string
	Distance    int
	GallonsUsed int
	CostOfFuel  float32
	StartDest   string
	EndDest     string
}
