package openapi

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func getVersionWithMocks() *http.Response {
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "/version", nil)
	GetVersion(recorder, request)
	return recorder.Result()
}

func TestGetVersionStatus(t *testing.T) {
	response := getVersionWithMocks()
	equals(t, http.StatusOK, response.StatusCode)
}

func TestGetVersionContentType(t *testing.T) {
	response := getVersionWithMocks()
	equals(t, "application/json; charset=UTF-8", response.Header.Get("Content-Type"))
}

func TestGetVersionBody(t *testing.T) {
	response := getVersionWithMocks()
	body, err := ioutil.ReadAll(response.Body)
	ok(t, err)

	actual := Version{}
	json.Unmarshal(body, &actual)
	equals(t, Version{APIVersion, BuildDate}, actual)
}
