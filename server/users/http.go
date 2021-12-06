package users

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"architecture-lab-3/server/tools"
)

// Channels HTTP handler.
type HttpHandlerFunc http.HandlerFunc

// HttpHandler creates a new instance of channels HTTP handler.
func HttpHandler(store *Store) HttpHandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			handleUserCreate(r, rw, store)
		} else {
			rw.WriteHeader(http.StatusMethodNotAllowed)
		}
	}
}

func handleUserCreate(r *http.Request, rw http.ResponseWriter, store *Store) {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		log.Printf("Error: %s", err)
		tools.WriteJsonBadRequest(rw, "Invalid JSON")
		return
	}

	if err := validateUser(user); err != nil {
		tools.WriteJsonBadRequest(rw, err.Error())
		return
	}

	err := store.CreateUser(&user)
	if err == nil {
		tools.WriteJsonOk(rw, &user)
	} else {
		log.Printf("Error inserting record: %s", err)
		tools.WriteJsonInternalError(rw)
	}
}

func validateUser(user User) error {
	if len(user.userName) == 0 {
		return errors.New("userName may not be empty")
	}
	return nil
}
