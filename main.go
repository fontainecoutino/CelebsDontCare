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
	// user.SetupRoutes(basePath)

	// start service :)
	fmt.Println("Started service ...")
	log.Fatal(http.ListenAndServe(":5000", nil))
}
