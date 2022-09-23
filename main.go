package main

import (
	"log"
	"net/http"

	"github.com/fontainecoutino/CelebsDontCare/database"
	"github.com/fontainecoutino/CelebsDontCare/trip"
	user "github.com/fontainecoutino/CelebsDontCare/users"
)

const basePath = "/api"

func main() {
	database.SetupDatabase()
	trip.SetupRoutes(basePath)
	user.SetupRoutes(basePath)
	log.Fatal(http.ListenAndServe(":5000", nil))
}
