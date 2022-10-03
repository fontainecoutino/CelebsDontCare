package trip

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/fontainecoutino/CelebsDontCare/cors"
)

const tripsPath = "trips"

// SetupRoutes
func SetupRoutes(apiBasePath string) {
	// /apiPath/trips
	tripsHandler := http.HandlerFunc(handleTrips)
	http.Handle(fmt.Sprintf("%s/%s", apiBasePath, tripsPath), cors.Middleware(tripsHandler))

	// /apiPath/trips/
	tripHandler := http.HandlerFunc(handleTrip)
	http.Handle(fmt.Sprintf("%s/%s/", apiBasePath, tripsPath), cors.Middleware(tripHandler))

	// /apiPath/trips/gets-data-display
	dataDisplayHandler := http.HandlerFunc(handleDataDisplay)
	http.Handle(fmt.Sprintf("%s/%s/gets-data-display", apiBasePath, tripsPath), cors.Middleware(dataDisplayHandler))
}

func handleTrip(w http.ResponseWriter, r *http.Request) {
	// get segments
	urlPathSegments := strings.Split(r.URL.Path, tripsPath+"/")
	if len(urlPathSegments[1:]) > 1 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// get tripID
	tripID, err := strconv.Atoi(urlPathSegments[len(urlPathSegments)-1])
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	switch r.Method {
	// GET
	case http.MethodGet:
		product, err := getTrip(tripID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if product == nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		j, err := json.Marshal(product)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		_, err = w.Write(j)
		if err != nil {
			log.Fatal(err)
		}

	// DELETE
	case http.MethodDelete:
		err := removeTrip(tripID)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

	case http.MethodOptions:
		return

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func handleTrips(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	// GET
	case http.MethodGet:
		tripList, err := getTripList()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		j, err := json.Marshal(tripList)
		if err != nil {
			log.Fatal(err)
		}
		_, err = w.Write(j)
		if err != nil {
			log.Fatal(err)
		}

	// POST
	case http.MethodPost:
		var trip Trip
		err := json.NewDecoder(r.Body).Decode(&trip)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err = insertTrip(trip)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusCreated)

	case http.MethodOptions:
		return

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func handleDataDisplay(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	// GET
	case http.MethodGet:
		dataDisplay, err := getDataDisplay()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		j, err := json.Marshal(dataDisplay)
		if err != nil {
			log.Fatal(err)
		}
		_, err = w.Write(j)
		if err != nil {
			log.Fatal(err)
		}

	case http.MethodOptions:
		return

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
