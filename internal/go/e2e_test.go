// +build large_test

package openapi

import (
	"os"
	"testing"

	"github.com/adammansfield/grocerer/pkg/ourgrocer"
)

const (
	emailEnvVar = "OURGROCERIES_EMAIL"
	passEnvVar  = "OURGROCERIES_PASSWORD"
)

func TestAddItem(t *testing.T) {
	email := os.Getenv(emailEnvVar)
	pass := os.Getenv(passEnvVar)
	assert(t, email != "", "Environment variable %s is empty or not set", emailEnvVar)
	assert(t, pass != "", "Environment variable %s is empty or not set", passEnvVar)

	client := ourgrocer.OGClient{}
	err := client.Login(email, pass)
	ok(t, err)
	assert(t, client.TeamID != "", "teamID not found")

	lists, err := client.GetLists()
	ok(t, err)
	assert(t, len(lists) > 0, "lists not found")

	listID := lists[0].ID
	err = client.AddItem(listID, "sardines")
	ok(t, err)
}
