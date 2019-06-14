package openapi

import (
	"encoding/json"
	"net/http"
)

const (
	VERSION    = "1"
	BUILD_DATE = "2019-06-14T12:23"
)

func GetVersion(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	v := Version{VERSION, BUILD_DATE}
	json.NewEncoder(w).Encode(v)
}
