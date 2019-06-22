// +build large_test

package openapi

import (
	"fmt"
	"testing"
)

func TestGetLists(t *testing.T) {
	assert(t, container.Config != Config{}, "empty config")

	client := OGClient{}
	err := client.Login()
	ok(t, err)
	assert(t, client.TeamID != "", "teamID not found")
	fmt.Printf("  teamID: %v\n", client.TeamID)

	lists, err := client.GetLists()
	ok(t, err)
	assert(t, len(lists) > 0, "lists not found")
	fmt.Printf("  lists: %+v\n", lists)
}
