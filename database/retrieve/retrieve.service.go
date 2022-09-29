package retrieve

import (
	"fmt"
	"net/http"

	"github.com/fontainecoutino/CelebsDontCare/cors"
)

const retrievePath = "retrieve"
const currentSource = "twitter"

// SetupRoutes
func SetupRoutes(apiBasePath string) {
	retrieveHandler := http.HandlerFunc(handleRetrieve)
	http.Handle(apiBasePath+"/"+retrievePath, cors.Middleware(retrieveHandler))
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
