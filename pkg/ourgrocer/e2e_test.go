// +build large_test

package ourgrocer_test

import (
	"encoding/base64"
	"math/rand"
	"os"
	"strings"
	"testing"

	"github.com/adammansfield/grocerer/pkg/ourgrocer"
)

// Test fixture for storing state between tests to avoid multiple logins
type fixture struct {
	Client  ourgrocer.Client
	Email   string             // Ourgroceries email
	Item    string             // Item to be added using AddItem
	ListIDs []ourgrocer.ListID // Lists returned by GetLists
	Pass    string             // Ourgroceries password
}

var f fixture

func TestSetup(t *testing.T) {
	f = newFixture(t)
}

func TestLogin(t *testing.T) {
	err := f.Client.Login(f.Email, f.Pass)
	ok(t, err)
	assert(t, f.Client.TeamID != "", "teamID not found")
}

func TestGetLists(t *testing.T) {
	var err error
	f.ListIDs, err = f.Client.GetLists()
	ok(t, err)
	assert(t, len(f.ListIDs) > 0, "lists not found")
}

func TestAddItem(t *testing.T) {
	err := f.Client.AddItem(f.listID(), f.Item)
	ok(t, err)
}

func TestGetList(t *testing.T) {
	items, err := f.Client.GetList(f.listID())
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

func getPass(t *testing.T) string {
	result := os.Getenv("OURGROCERIES_PASSWORD")
	assert(t, result != "", "Environment variable %s is empty or not set", "OURGROCERIES_PASSWORD")
	return result
}

func (f *fixture) listID() string {
	return f.ListIDs[0].ID
}

func newFixture(t *testing.T) fixture {
	f := fixture{}
	f.Client = ourgrocer.Client{}
	f.Email = getEmail(t)
	f.Item = generateItem(t)
	f.Pass = getPass(t)
	return f
}
