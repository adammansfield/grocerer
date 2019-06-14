package openapi

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func GetVersionWithMocks() *http.Response {
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "/version", nil)
	GetVersion(recorder, request)
	return recorder.Result()
}

func TestGetVersionStatus(t *testing.T) {
	response := GetVersionWithMocks()
	expected := http.StatusOK
	actual := response.StatusCode
	if expected != actual {
		t.Errorf("expected: %+v, actual: %+v", expected, actual)
	}
}

func TestGetVersionContentType(t *testing.T) {
	response := GetVersionWithMocks()
	expected := "application/json; charset=UTF-8"
	actual := response.Header.Get("Content-Type")
	if expected != actual {
		t.Errorf("expected: %+v, actual: %+v", expected, actual)
	}
}

func TestGetVersionBody(t *testing.T) {
	response := GetVersionWithMocks()
	body, _ := ioutil.ReadAll(response.Body)

	expected := Version{APIVersion, BuildDate}
	actual := &Version{}
	json.Unmarshal(body, actual)
	if expected != *actual {
		t.Errorf("expected: %+v, actual: %+v", expected, *actual)
	}
}
