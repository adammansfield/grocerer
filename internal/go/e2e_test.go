// +build large_test

package openapi

import (
	"os"
	"testing"
)

const (
	emailEnvVar = "OURGROCERIES_EMAIL"
	passEnvVar  = "OURGROCERIES_PASSWORD"
)

func TestAddItem(t *testing.T) {
	email := os.Getenv(emailEnvVar)
	pass := os.Getenv(passEnvVar)
	assert(t, email != "", "Enviornment variable %s is empty or not set", emailEnvVar)
	assert(t, pass != "", "Enviornment variable %s is empty or not set", passEnvVar)

	client := OGClient{}
	err := client.Login(email, pass)
	ok(t, err)
	assert(t, client.TeamID != "", "teamID not found")

	lists, err := client.GetLists()
	ok(t, err)
	assert(t, len(lists) > 0, "lists not found")

	listID := lists[0].Id
	err = client.AddItem(listID, "sardines")
	ok(t, err)
}
