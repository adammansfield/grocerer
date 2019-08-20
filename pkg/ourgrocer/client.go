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
	teamIDRegEx = regexp.MustCompile(`g_teamId = "([A-Za-z0-9]*)"`)
)

// Client is an OurGroceries client.
type Client struct {
	cookieJar  *cookiejar.Jar
	httpClient http.Client
	teamID     string
}

// NewClient returns a Client.
func NewClient(cookieJar *cookiejar.Jar, httpClient http.Client) Client {
	result := Client{}
	result.cookieJar = cookieJar
	result.httpClient = httpClient
	return result
}

// AddItem adds an item to the given list.
func (client *Client) AddItem(listID string, item string) error {
	_, err := client.call(command{"insertItem", client.teamID, listID, item})
	return err
}

func (client *Client) call(cmd command) ([]byte, error) {
	request, err := buildYourListsRequest(cmd)
	if err != nil {
		return nil, err
	}

	response, err := client.httpClient.Do(request)
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

// GetList gets grocery items for a list.
func (client *Client) GetList(listID string) ([]Item, error) {
	body, err := client.call(command{"getList", client.teamID, listID, ""})
	return HandleGetList(body, err)
}

// GetLists gets the grocery lists from OurGroceries.
func (client *Client) GetLists() ([]ListID, error) {
	body, err := client.call(command{"getOverview", client.teamID, "", ""})
	return handleGetLists(body, err)
}

// Login authenticates with OurGroceries and returns the user's teamID
func (client *Client) Login(email string, password string) error {
	request, err := buildLoginRequest(email, password)
	if err != nil {
		return err
	}

	response, err := client.httpClient.Do(request)
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
	cookies := client.cookieJar.Cookies(signInURL)
	if cookies == nil {
		return fmt.Errorf("invalid credentials")
	}

	client.teamID, err = ExtractTeamID(response.Body)
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

// ExtractTeamID returns teamID from the response body of /sign-in.
// TODO: make ExtractTeamID a private function
func ExtractTeamID(r io.Reader) (string, error) {
	bytes, err := ioutil.ReadAll(r)
	if err != nil {
		return "", err
	}

	teamIds := teamIDRegEx.FindStringSubmatch(string(bytes))
	if len(teamIds) < 2 {
		return "", fmt.Errorf("teamID not found in body")
	}
	return teamIds[1], nil
}

// HandleGetList returns []Item from the response body of getList.
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
