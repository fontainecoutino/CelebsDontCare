package retrieve

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/fontainecoutino/CelebsDontCare/cors"
)

const retrievePath = "retrieve"
const currentSource = "twitter"

// SetupRoutes
func SetupRoutes(apiBasePath string) {
	retrieveHandler := http.HandlerFunc(handleRetrieve)
	retrieveHandlerPath := apiBasePath + "/" + retrievePath
	http.Handle(retrieveHandlerPath, cors.Middleware(retrieveHandler))

	retrieveSpecificHandler := http.HandlerFunc(handleRetreiveSpecific)
	retrieveSpecificHandlerPath := apiBasePath + "/" + retrievePath + "/"
	http.Handle(retrieveSpecificHandlerPath, cors.Middleware(retrieveSpecificHandler))
}

func handleRetrieve(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	// GET
	case http.MethodGet:
		newTrips, err := getData(currentSource)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(fmt.Sprintf(`{"num_new_trips":%d}`, newTrips)))

	case http.MethodOptions:
		return

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func handleRetreiveSpecific(w http.ResponseWriter, r *http.Request) {
	// get segments
	urlPathSegments := strings.Split(r.URL.Path, retrievePath+"/")
	if len(urlPathSegments[1:]) > 1 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// get tripID
	source := urlPathSegments[len(urlPathSegments)-1]

	switch r.Method {
	// GET
	case http.MethodGet:
		newTrips, err := getData(source)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(fmt.Sprintf(`{"productId":%d}`, newTrips)))

	case http.MethodOptions:
		return

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
