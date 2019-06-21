package openapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strings"
)

const (
	mainRawURL      = "https://www.ourgroceries.com"
	signInRawURL    = "https://www.ourgroceries.com/sign-in"
	yourListsRawURL = "https://www.ourgroceries.com/your-lists/"
	userAgent       = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.90 Safari/537.36"
)

// login authenticates with OurGroceries and returns the user's teamId
func login() (string, error) {
	form := url.Values{}
	form.Set("action", "sign-me-in")
	form.Set("emailAddress", container.Config.Email)
	form.Set("password", container.Config.Password)
	form.Set("staySignedIn", "on")

	request, err := http.NewRequest("POST", signInRawURL, strings.NewReader(form.Encode()))
	if err != nil {
		return "", err
	}
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Add("Origin", mainRawURL)
	request.Header.Add("Referer", signInRawURL)
	request.Header.Add("User-Agent", userAgent)

	response, err := container.HTTPClient.Do(request)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code %d for %s", response.StatusCode, response.Request.URL.String())
	}

	signInURL, err := url.Parse(signInRawURL)
	if err != nil {
		return "", err
	}
	cookies := container.CookieJar.Cookies(signInURL)
	if cookies == nil {
		return "", fmt.Errorf("invalid credentials")
	}

	return extractTeamID(response.Body)
}

// extractTeamId returns teamId from the response body of /sign-in
func extractTeamID(stream io.Reader) (string, error) {
	bytes, err := ioutil.ReadAll(stream)
	if err != nil {
		return "", err
	}

	re := regexp.MustCompile(`g_teamId = "([A-Za-z0-9]*)"`)
	teamIds := re.FindStringSubmatch(string(bytes))
	if len(teamIds) < 2 {
		return "", fmt.Errorf("teamId not found in body")
	}
	return teamIds[1], nil
}

type command struct {
	Command string `json:"command"`
	TeamID  string `json:"teamId"`
}

type yourListsResponse struct {
	ShoppingLists []List `json:"shoppingLists"`
}

// getLists gets the grocery lists from OurGroceries
func getLists(teamID string) ([]List, error) {
	getListsCommand := command{"getOverview", teamID}
	body, err := json.Marshal(getListsCommand)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", yourListsRawURL, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	request.Header.Add("Accept", "application/json, text/javascript, */*")
	request.Header.Add("Content-Type", "application/json; charset=UTF-8")
	request.Header.Add("Host", "www.ourgroceries.com")
	request.Header.Add("Origin", mainRawURL)
	request.Header.Add("Referer", yourListsRawURL)
	request.Header.Add("User-Agent", userAgent)
	request.Header.Add("X-Requested-With", "XMLHttpRequest")

	response, err := container.HTTPClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code %d for %s", response.StatusCode, response.Request.URL.String())
	}

	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	if response.Header.Get("Content-Type") != "application/json; charset=UTF-8" {
		return nil, fmt.Errorf("unxpected content type %s for %s", response.Header.Get("Content-Type"), response.Request.URL.String())
	}

	yourListsResponseJSON := yourListsResponse{}
	err = json.Unmarshal(bytes, &yourListsResponseJSON)
	if err != nil {
		return nil, err
	}

	return yourListsResponseJSON.ShoppingLists, nil
}

// GetLists responsd with the grocery lists
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
