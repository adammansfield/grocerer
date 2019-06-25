package openapi

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// AddItem adds an item to a list
func AddItem(w http.ResponseWriter, r *http.Request) {
	client := OGClient{}
	if client.Login() != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	item := Item{}
	if json.NewDecoder(r.Body).Decode(&item) != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	lists, err := client.GetLists()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	args := mux.Vars(r)
	listID, err := strconv.Atoi(args["listID"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if client.AddItem(lists[listID].Id, item.Name) != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
}

// GetLists responds with the grocery lists
func GetLists(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

// GetVersion responds with the API version and build date
func GetVersion(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(PackageVersion)
}
