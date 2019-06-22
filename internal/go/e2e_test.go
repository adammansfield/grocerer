// +build large_test

package openapi

import (
	"fmt"
	"testing"
)

func TestAddItem(t *testing.T) {
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

	listID := lists[0].Id
	err = client.AddItem(listID, "baby dill pickles")
	ok(t, err)
	fmt.Printf("  itemAdded: %v\n", "baby dill pickles")
}
