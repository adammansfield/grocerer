package openapi

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/adammansfield/grocerer/pkg/ourgrocer"
	"github.com/gorilla/mux"
)

// AddItem adds an item to a list
func AddItem(w http.ResponseWriter, r *http.Request) {
	client, err := login(r)
	if err != nil {
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

	if client.AddItem(lists[listID].ID, item.Name) != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
}

// GetList responds with a grocery list
func GetList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

// GetLists responds with the grocery lists
func GetLists(w http.ResponseWriter, r *http.Request) {
	client, err := login(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	lists, err := client.GetLists()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(lists)
}

// GetVersion responds with the API version and build date
func GetVersion(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(PackageVersion)
}

func login(r *http.Request) (ourgrocer.Client, error) {
	email := r.URL.Query().Get("email")
	password := r.URL.Query().Get("password")

	client := ourgrocer.Client{}
	err := client.Login(email, password)
	return client, err
}
