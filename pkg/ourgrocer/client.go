package ourgrocer

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
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

var (
	re           = regexp.MustCompile(`g_teamId = "([A-Za-z0-9]*)"`)
	cookieJar, _ = cookiejar.New(nil)
	httpClient   = http.Client{Jar: cookieJar}
)

// Client is an OurGroceries client
type Client struct {
	TeamID string
}

// AddItem adds an item to the given list
func (client *Client) AddItem(listID string, item string) error {
	_, err := call(command{"insertItem", client.TeamID, listID, item})
	return err
}

// GetList gets grocery items for a list
func (client *Client) GetList(listID string) ([]Item, error) {
	body, err := call(command{"getList", client.TeamID, listID, ""})
	return HandleGetList(body, err)
}

// GetLists gets the grocery lists from OurGroceries
func (client *Client) GetLists() ([]ListID, error) {
	body, err := call(command{"getOverview", client.TeamID, "", ""})
	return handleGetLists(body, err)
}

// Login authenticates with OurGroceries and returns the user's teamId
func (client *Client) Login(email string, password string) error {
	request, err := buildLoginRequest(email, password)
	if err != nil {
		return err
	}

	response, err := httpClient.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	err = handleStatusCode(response)
	if err != nil {
		return err
	}

	signInURL, err := url.Parse(signInRawURL)
	if err != nil {
		return err
	}
	cookies := cookieJar.Cookies(signInURL)
	if cookies == nil {
		return fmt.Errorf("invalid credentials")
	}

	client.TeamID, err = ExtractTeamID(response.Body)
	return err
}

func buildYourListsRequest(cmd command) (*http.Request, error) {
	body, err := json.Marshal(cmd)
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
	return request, nil
}

func buildLoginRequest(email string, password string) (*http.Request, error) {
	form := url.Values{}
	form.Set("action", "sign-me-in")
	form.Set("emailAddress", email)
	form.Set("password", password)
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

func call(cmd command) ([]byte, error) {
	request, err := buildYourListsRequest(cmd)
	if err != nil {
		return nil, err
	}

	response, err := httpClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	err = handleStatusCode(response)
	if err != nil {
		return nil, err
	}

	return ioutil.ReadAll(response.Body)
}

// ExtractTeamID returns teamId from the response body of /sign-in
// TODO: make ExtractTeamID a private function
func ExtractTeamID(r io.Reader) (string, error) {
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

// HandleGetList returns []Item from the response body of getList
// TODO: make HandleGetList a private function
func HandleGetList(bytes []byte, err error) ([]Item, error) {
	if err != nil {
		return nil, err
	}

	result := getListResponse{}
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		return nil, err
	}

	return result.List.Items, err
}

func handleGetLists(bytes []byte, err error) ([]ListID, error) {
	if err != nil {
		return nil, err
	}

	response := getOverviewResponse{}
	err = json.Unmarshal(bytes, &response)
	if err != nil {
		return nil, err
	}

	return response.ShoppingLists, nil
}

func handleStatusCode(response *http.Response) error {
	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code %d for %s", response.StatusCode, response.Request.URL.String())
	}

	return nil
}
