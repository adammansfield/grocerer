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
	host            = "www.ourgroceries.com"
	mainRawURL      = "https://www.ourgroceries.com"
	signInRawURL    = "https://www.ourgroceries.com/sign-in"
	yourListsRawURL = "https://www.ourgroceries.com/your-lists/"
	userAgent       = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.90 Safari/537.36"
)

var re = regexp.MustCompile(`g_teamId = "([A-Za-z0-9]*)"`)

// OGClient is a client for OurGroceries
type OGClient struct {
	TeamID string
}

type command struct {
	Command string `json:"command"`
	ListID  string `json:"listId,omitempty"`
	TeamID  string `json:"teamId"`
	Value   string `json:"value,omitempty"`
}

type yourListsResponse struct {
	ShoppingLists []List `json:"shoppingLists"`
}

func buildAddItemRequest(teamID string, listID string, item string) (*http.Request, error) {
	body, err := json.Marshal(command{Command: "insertItem", ListID: listID, TeamID: teamID, Value: item})
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", yourListsRawURL, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	request.Header.Add("Accept", "application/json, text/javascript, */*")
	request.Header.Add("Content-Type", "application/json; charset=UTF-8")
	request.Header.Add("Host", host)
	request.Header.Add("Origin", mainRawURL)
	request.Header.Add("Referer", yourListsRawURL)
	request.Header.Add("User-Agent", userAgent)
	request.Header.Add("X-Requested-With", "XMLHttpRequest")
	return request, nil
}

func buildLoginRequest() (*http.Request, error) {
	form := url.Values{}
	form.Set("action", "sign-me-in")
	form.Set("emailAddress", container.Config.Email)
	form.Set("password", container.Config.Password)
	form.Set("staySignedIn", "on")

	request, err := http.NewRequest("POST", signInRawURL, strings.NewReader(form.Encode()))
	if err != nil {
		return nil, err
	}

	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Add("Origin", mainRawURL)
	request.Header.Add("Referer", signInRawURL)
	request.Header.Add("User-Agent", userAgent)
	return request, nil
}

// buildListsRequest returns request for a secret get lists in JSON
func buildListsRequest(teamID string) (*http.Request, error) {
	body, err := json.Marshal(command{Command: "getOverview", TeamID: teamID})
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", yourListsRawURL, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	request.Header.Add("Accept", "application/json, text/javascript, */*")
	request.Header.Add("Content-Type", "application/json; charset=UTF-8")
	request.Header.Add("Host", host)
	request.Header.Add("Origin", mainRawURL)
	request.Header.Add("Referer", yourListsRawURL)
	request.Header.Add("User-Agent", userAgent)
	request.Header.Add("X-Requested-With", "XMLHttpRequest")
	return request, nil
}

// extractLists returns grocery lists from the response body of getOverview
func extractLists(body io.Reader) ([]List, error) {
	bytes, err := ioutil.ReadAll(body)
	if err != nil {
		return nil, err
	}

	response := yourListsResponse{}
	err = json.Unmarshal(bytes, &response)
	if err != nil {
		return nil, err
	}

	return response.ShoppingLists, nil
}

// extractTeamId returns teamId from the response body of /sign-in
func extractTeamID(r io.Reader) (string, error) {
	bytes, err := ioutil.ReadAll(r)
	if err != nil {
		return "", err
	}

	teamIds := re.FindStringSubmatch(string(bytes))
	if len(teamIds) < 2 {
		return "", fmt.Errorf("teamId not found in body")
	}
	return teamIds[1], nil
}

// GetLists gets the grocery lists from OurGroceries
func (client *OGClient) GetLists() ([]List, error) {
	request, err := buildListsRequest(client.TeamID)
	if err != nil {
		return nil, err
	}

	response, err := container.HTTPClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code %d for %s", response.StatusCode, response.Request.URL.String())
	}

	return extractLists(response.Body)
}

// Login authenticates with OurGroceries and returns the user's teamId
func (client *OGClient) Login() error {
	request, err := buildLoginRequest()
	if err != nil {
		return err
	}

	response, err := container.HTTPClient.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code %d for %s", response.StatusCode, response.Request.URL.String())
	}

	signInURL, err := url.Parse(signInRawURL)
	if err != nil {
		return err
	}
	cookies := container.CookieJar.Cookies(signInURL)
	if cookies == nil {
		return fmt.Errorf("invalid credentials")
	}

	client.TeamID, err = extractTeamID(response.Body)
	return err
}

// AddItem adds an item to the given list
func (client *OGClient) AddItem(listID string, item string) error {
	request, err := buildAddItemRequest(client.TeamID, listID, item)
	if err != nil {
		return err
	}

	response, err := container.HTTPClient.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code %d for %s", response.StatusCode, response.Request.URL.String())
	}

	return nil
}
