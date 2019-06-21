// +build large_test

package openapi

import (
	"fmt"
	"testing"
)

func TestGetLists(t *testing.T) {
	assert(t, container.Config != Config{}, "empty config")

	teamID, err := login()
	ok(t, err)
	assert(t, teamID != "", "teamID not found")
	fmt.Printf("teamID: %v\n", teamID)

	lists, err := getLists(teamID)
	ok(t, err)
	assert(t, len(lists) > 0, "lists not found")
	fmt.Printf("lists: %+v\n", lists)
}
