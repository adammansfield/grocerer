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

func TestLogin(t *testing.T) {
	cookieJar, _ := cookiejar.New(nil)
	uri, _ := url.Parse("https://www.ourgroceries.com/sign-in")
	cookieJar.SetCookies(uri, []*http.Cookie{{}})

	httpClient := httpClientMock{}
	httpClient.err = nil
	httpClient.response = &http.Response{}
	httpClient.response.Body = ioutil.NopCloser(strings.NewReader("var g_teamId = \"E0KAegvBF9SOQ78b9vhlYr\""))
	httpClient.response.StatusCode = http.StatusOK

	client := ourgrocer.NewClient(cookieJar, &httpClient)
	ok(t, client.Login("email", "pass"))
}

// httpClientMock implements ourgrocer.HTTPClient.
type httpClientMock struct {
	response *http.Response
	err      error
}

func (client *httpClientMock) Do(request *http.Request) (*http.Response, error) {
	return client.response, client.err
}
