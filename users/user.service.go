package user

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/fontainecoutino/CelebsDontCare/cors"
)

const usersPath = "users"

// SetupRoutes
func SetupRoutes(apiBasePath string) {
	userHandler := http.HandlerFunc(handleUser)
	userHandlerPath := apiBasePath + "/" + usersPath + "/"
	http.Handle(userHandlerPath, cors.Middleware(userHandler))

	usersHandler := http.HandlerFunc(handleUsers)
	usersHandlerPath := apiBasePath + "/" + usersPath
	http.Handle(usersHandlerPath, cors.Middleware(usersHandler))
}

func handleUser(w http.ResponseWriter, r *http.Request) {
	// get segments
	urlPathSegments := strings.Split(r.URL.Path, usersPath+"/")
	if len(urlPathSegments[1:]) > 1 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// get userID
	userID, err := strconv.Atoi(urlPathSegments[len(urlPathSegments)-1])
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	switch r.Method {
	// GET
	case http.MethodGet:
		product, err := getUser(userID)
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
		err := removeUser(userID)
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

func handleUsers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	// GET
	case http.MethodGet:
		userList, err := getUserList()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		j, err := json.Marshal(userList)
		if err != nil {
			log.Fatal(err)
		}
		_, err = w.Write(j)
		if err != nil {
			log.Fatal(err)
		}

	// POST
	case http.MethodPost:
		var user User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err = insertUser(user)
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
