package main

import (
	"log"
	"net/http"

	"github.com/fontainecoutino/CelebsDontCare/database"
	"github.com/fontainecoutino/CelebsDontCare/trip"
)

const basePath = "/api"

type Trip struct {
	TripID      int
	TimeStamp   string
	UserID      string
	Distance    int
	GallonsUsed int
	CostOfFuel  float32
	StartDest   string
	EndDest     string
}

func main() {
	database.SetupDatabase()
	trip.SetupRoutes(basePath)
	log.Fatal(http.ListenAndServe(":5000", nil))

}
