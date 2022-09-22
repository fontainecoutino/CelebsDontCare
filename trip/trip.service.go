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

// SetupRoutes
func SetupRoutes(apiBasePath string) {
	tripHandler := http.HandlerFunc(handleTrip)
	tripHandlerPath := apiBasePath + "/" + tripsPath + "/"

	//tripsHandler := http.HandlerFunc(handleProducts)
	//tripsHandlerPath := apiBasePath + "/" + tripsPath

	http.Handle(tripHandlerPath, cors.Middleware(tripHandler))
	//http.Handle(tripsHandlerPath, cors.Middleware(tripsHandler))
}
