package trip

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/fontainecoutino/CelebsDontCare/cors"
)

const tripsPath = "trips"

// SetupRoutes
func SetupRoutes(apiBasePath string) {
	tripHandler := http.HandlerFunc(handleTrip)
	tripHandlerPath := apiBasePath + "/" + tripsPath + "/"

	tripsHandler := http.HandlerFunc(handleTrips)
	tripsHandlerPath := apiBasePath + "/" + tripsPath

	http.Handle(tripHandlerPath, cors.Middleware(tripHandler))
	http.Handle(tripsHandlerPath, cors.Middleware(tripsHandler))
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
		product, err := getProduct(tripID)
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

		/*
			// PUT
			case http.MethodPut:
				var product Product
				err := json.NewDecoder(r.Body).Decode(&product)
				if err != nil {
					log.Print(err)
					w.WriteHeader(http.StatusBadRequest)
					return
				}
				if *product.ProductID != productID {
					w.WriteHeader(http.StatusBadRequest)
					return
				}
				err = updateProduct(product)
				if err != nil {
					log.Print(err)
					w.WriteHeader(http.StatusBadRequest)
					return
				}

			// DELETE
			case http.MethodDelete:
				err := removeProduct(productID)
				if err != nil {
					log.Print(err)
					w.WriteHeader(http.StatusInternalServerError)
					return
				}
		*/
	case http.MethodOptions:
		return

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func handleTrips(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	/*
		// GET
		case http.MethodGet:
			productList, err := getProductList()
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			j, err := json.Marshal(productList)
			if err != nil {
				log.Fatal(err)
			}
			_, err = w.Write(j)
			if err != nil {
				log.Fatal(err)
			}
	*/
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
