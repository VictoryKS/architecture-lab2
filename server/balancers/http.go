package balancers

import (
	"encoding/json"
	"github.com/rebase/architecture-lab2/server/tools"
	"log"
	"net/http"
	"strconv"
)

// Balancers HTTP handler.
type HttpHandlerFunc http.HandlerFunc

// HttpHandler creates a new instance of balancers HTTP handler.
func HttpHandler(store *Store) HttpHandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			handleGetBalancer(store, rw, r)
		} else if r.Method == "POST" {
			handleMachineStatus(r, rw, store)
		} else {
			rw.WriteHeader(http.StatusMethodNotAllowed)
		}
	}
}

func handleMachineStatus(r *http.Request, rw http.ResponseWriter, store *Store) {
	var m Machine
	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		log.Printf("Error decoding machine input: %s", err)
		tools.WriteJsonBadRequest(rw, "bad JSON payload")
		return
	}
	err := store.MachineStatus(m.Id, m.IsWorking)
	if err == nil {
		tools.WriteJsonOk(rw, &m)
	} else {
		log.Printf("Error inserting record: %s", err)
		tools.WriteJsonInternalError(rw)
	}
}

func handleGetBalancer(store *Store, rw http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		log.Printf("Wrong URL: %s", err)
		tools.WriteJsonInternalError(rw)
		return
	}
	res, err := store.GetBalancer(id)
	if err != nil {
		log.Printf("Error making query to the db: %s", err)
		tools.WriteJsonInternalError(rw)
		return
	}
	tools.WriteJsonOk(rw, res)
}
