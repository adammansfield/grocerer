package openapi

import (
	"encoding/json"
	"net/http"
)

const (
	// APIVersion is the version of the API
	APIVersion = "1"
	// BuildDate is the date of the build
	BuildDate = "2019-06-14T12:23"
)

// GetVersion responds with the API version and build date
func GetVersion(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	v := Version{APIVersion, BuildDate}
	json.NewEncoder(w).Encode(v)
}
