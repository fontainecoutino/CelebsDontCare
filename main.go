package main

import (
	"log"
	"net/http"

	"github.com/fontainecoutino/CelebsDontCare/database"
	"github.com/fontainecoutino/CelebsDontCare/trip"
)

const basePath = "/api"

func main() {
	database.SetupDatabase()
	trip.SetupRoutes(basePath)
	log.Fatal(http.ListenAndServe(":5000", nil))
}
