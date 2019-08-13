// +build small_test

package ourgrocer_test

import (
	"strings"
	"testing"

	"github.com/adammansfield/grocerer/pkg/ourgrocer"
)

// TODO: modify test to only call Login (and still test teamID extraction)
func TestExtractTeamId(t *testing.T) {
	stream := strings.NewReader("var g_teamId = \"E0KAegvBF9SOQ78b9vhlYr\"")
	teamID, err := ourgrocer.ExtractTeamID(stream)
	ok(t, err)
	equals(t, "E0KAegvBF9SOQ78b9vhlYr", teamID)
}

// TODO: modify test to only call GetList (and still test grocery list parsing)
func TestParseGroceryList(t *testing.T) {
	json := `
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
	}`
	stream := strings.NewReader(json)
	items, err := ourgrocer.ParseGroceryList(stream)
	ok(t, err)
	equals(t, items, []ourgrocer.Item{{ID: "VVbCucm4eT30FIW9ptejCr", Value: "celery", CategoryID: "ow7os337oMPoE2RnZPelRI"}, {ID: "irezU2ekUw34Sbk8YoNG3Q", Value: "cherries"}})
}
