package openapi

import (
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
	mainRawURL   = "https://www.ourgroceries.com"
	signInRawURL = "https://www.ourgroceries.com/sign-in"
	userAgent    = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.90 Safari/537.36"
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

// GetVersion responds with the API version and build date
func GetVersion(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(PackageVersion)
}
