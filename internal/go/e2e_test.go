// +build large_test

package openapi

import (
  "testing"
)

func TestLogIn(t *testing.T) {
  assert(t, container.Config != Config{}, "empty config")
  ok(t, logIn())
}
