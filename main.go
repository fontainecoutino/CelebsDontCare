package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/fontainecoutino/CelebsDontCare/database"
	"github.com/fontainecoutino/CelebsDontCare/database/retrieve"
	"github.com/fontainecoutino/CelebsDontCare/templates"
	"github.com/fontainecoutino/CelebsDontCare/trip"
)

const basePath = "/api"

func main() {
	// front end
	templates.SetupRoutes()

	// back end
	database.SetupDatabase()
	trip.SetupRoutes(basePath)
	retrieve.SetupRoutes(basePath)
	// user.SetupRoutes(basePath) // not currently being used

	// sets up data retrieval for every hour

	// start service :)
	fmt.Println("Started service on http://localhost:4200/ ...")
	log.Fatal(http.ListenAndServe(":4200", nil))
}
