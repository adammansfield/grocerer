// +build large_test

package ourgrocer_test

import (
	"encoding/base64"
	"math/rand"
	"net/http"
	"net/http/cookiejar"
	"os"
	"strings"
	"testing"

	"github.com/adammansfield/grocerer/pkg/ourgrocer"
)

// Test fixture for storing state between tests to avoid multiple logins
type fixture struct {
	Client   ourgrocer.Client
	Email    string             // Ourgroceries email
	Item     string             // Item to be added using AddItem
	ListIDs  []ourgrocer.ListID // Lists returned by GetLists
	Password string             // Ourgroceries password
}

var f fixture

func TestSetup_e2e(t *testing.T) {
	f = newFixture(t)
}

func TestLogin_e2e(t *testing.T) {
	err := f.Client.Login(f.Email, f.Password)
	ok(t, err)
}

func TestGetLists_e2e(t *testing.T) {
	var err error
	f.ListIDs, err = f.Client.GetLists()
	ok(t, err)
	assert(t, len(f.ListIDs) > 0, "lists not found")
}

func TestAddItem_e2e(t *testing.T) {
	err := f.Client.AddItem(f.listID(t), f.Item)
	ok(t, err)
}

func TestGetList_e2e(t *testing.T) {
	items, err := f.Client.GetList(f.listID(t))
	ok(t, err)
	assert(t, containsName(items, f.Item), "%s was not added to the grocery list", f.Item)
}

func containsName(items []ourgrocer.Item, name string) bool {
	for _, item := range items {
		if strings.Contains(item.Value, name) {
			return true
		}
	}
	return false
}

func generateItem(t *testing.T) string {
	buf := make([]byte, 24)
	_, err := rand.Read(buf)
	ok(t, err)
	return base64.StdEncoding.EncodeToString(buf)
}

func getEmail(t *testing.T) string {
	result := os.Getenv("OURGROCERIES_EMAIL")
	assert(t, result != "", "Environment variable %s is empty or not set", "OURGROCERIES_EMAIL")
	return result
}

func getPassword(t *testing.T) string {
	result := os.Getenv("OURGROCERIES_PASSWORD")
	assert(t, result != "", "Environment variable %s is empty or not set", "OURGROCERIES_PASSWORD")
	return result
}

func (f *fixture) listID(t *testing.T) string {
	assert(t, len(f.ListIDs) > 0, "failed to get lists in previous a step")
	return f.ListIDs[0].ID
}

func newClient() ourgrocer.Client {
	cookieJar, _ := cookiejar.New(nil)
	httpClient := http.Client{Jar: cookieJar}
	return ourgrocer.NewClient(cookieJar, &httpClient)
}

func newFixture(t *testing.T) fixture {
	f := fixture{}
	f.Client = newClient()
	f.Email = getEmail(t)
	f.Item = generateItem(t)
	f.Password = getPassword(t)
	return f
}
