// +build small_test

package ourgrocer_test

import (
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
	"testing"

	"github.com/adammansfield/grocerer/pkg/ourgrocer"
)

// TODO: modify test to only call GetList (and still test grocery list parsing)
func TestHandleGetList(t *testing.T) {
	body := []byte(`
	{
	  "list":{
	    "notesHtml":null,
	    "versionId":"p3TC4L9k6wiqYUvB6Lvxrl",
	    "notes":null,
	    "name":"Groceries",
	    "id":"n8w1aAoqqURBrM2nUns6nQ",
	    "listType":"SHOPPING",
	    "items":[
	      {
	        "id":"VVbCucm4eT30FIW9ptejCr",
	        "value":"celery",
	        "categoryId":"ow7os337oMPoE2RnZPelRI"
	      },
	      {
	        "id":"irezU2ekUw34Sbk8YoNG3Q",
	        "value":"cherries"
	      }
	    ]
	  }
	}`)
	items, err := ourgrocer.HandleGetList(body, nil)
	ok(t, err)
	equals(t, items, []ourgrocer.Item{{ID: "VVbCucm4eT30FIW9ptejCr", Value: "celery", CategoryID: "ow7os337oMPoE2RnZPelRI"}, {ID: "irezU2ekUw34Sbk8YoNG3Q", Value: "cherries"}})
}

func TestAddItem(t *testing.T) {
	client, _, _ := newClientWithStubs()
	ok(t, client.AddItem("n8w1aAoqqURBrM2nUns6nQ", "grapes"))
}

func TestGetList(t *testing.T) {
	client, httpClient, _ := newClientWithStubs()
	httpClient.response.Body = ioutil.NopCloser(strings.NewReader("{\"list\":{\"items\":[{\"value\":\"Apples\"},{\"value\":\"Grapes\"}]}}"))
	items, err := client.GetList("n8w1aAoqqURBrM2nUns6nQ")
	ok(t, err)
	equals(t, items, []ourgrocer.Item{{Value: "Apples"}, {Value: "Grapes"}})
}

func TestGetLists(t *testing.T) {
	client, httpClient, _ := newClientWithStubs()
	httpClient.response.Body = ioutil.NopCloser(strings.NewReader("{\"shoppingLists\": [{\"name\": \"Groceries\"}]}"))
	listIDs, err := client.GetLists()
	ok(t, err)
	equals(t, listIDs, []ourgrocer.ListID{{Name: "Groceries"}})
}

func TestLogin(t *testing.T) {
	client, httpClient, _ := newClientWithStubs()
	httpClient.response.Body = ioutil.NopCloser(strings.NewReader("var g_teamId = \"E0KAegvBF9SOQ78b9vhlYr\""))
	ok(t, client.Login("email", "pass"))
}

func TestLoginCannotParseTeamID(t *testing.T) {
	client, _, _ := newClientWithStubs()
	assert(t, client.Login("email", "pass") != nil, "Invalid response body parsed")
}

func TestLoginInvalidCredentials(t *testing.T) {
	client, _, jar := newClientWithStubs()
	uri, _ := url.Parse("https://www.ourgroceries.com/sign-in")
	jar.SetCookies(uri, nil)
	assert(t, client.Login("email", "pass") != nil, "Invalid login accepted")
}

// httpClientStub implements ourgrocer.HTTPClient.
type httpClientStub struct {
	response *http.Response
	err      error
}

func (client *httpClientStub) Do(request *http.Request) (*http.Response, error) {
	return client.response, client.err
}

func newClientWithStubs() (ourgrocer.Client, *httpClientStub, *cookiejar.Jar) {
	jar, _ := cookiejar.New(nil)
	uri, _ := url.Parse("https://www.ourgroceries.com/sign-in")
	jar.SetCookies(uri, []*http.Cookie{{Name: "ourgroceries-auth", Value: "JbwTzQVF2ucvn0kufJE6Oc|16fe2ba873ce5879c934137a869568004fdd8ceabf32963b182437072214a8ee"}})

	httpClient := &httpClientStub{}
	httpClient.err = nil
	httpClient.response = &http.Response{}
	httpClient.response.Body = ioutil.NopCloser(strings.NewReader(""))
	httpClient.response.StatusCode = http.StatusOK

	client := ourgrocer.NewClient(jar, httpClient)
	return client, httpClient, jar
}
