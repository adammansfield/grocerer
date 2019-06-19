// +build small_test

package openapi

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func getVersionWithMocks() *http.Response {
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "/version", nil)
	GetVersion(recorder, request)
	return recorder.Result()
}

func TestExtractTeamId(t *testing.T) {
	stream := strings.NewReader("var g_teamId = \"E0KAegvBF9SOQ78b9vhlYr\"")
	teamID, err := extractTeamID(stream)
	ok(t, err)
	equals(t, "E0KAegvBF9SOQ78b9vhlYr", teamID)
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
	equals(t, PackageVersion, actual)
}
