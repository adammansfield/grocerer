// +build large_test

package openapi

import (
	"testing"
)

func TestLogIn(t *testing.T) {
	assert(t, container.Config != Config{}, "empty config")

	teamId, err := login()
	ok(t, err)
	assert(t, teamId != "", "teamId not found")
}
