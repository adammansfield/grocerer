package openapi

import (
	"encoding/json"
	"strings"
	"testing"
)

func createConfig(t *testing.T) (Config, Config) {
	expected := Config{"groceries@mail.test", "p@ssword"}
	json, err := json.Marshal(&expected)
	ok(t, err)

	reader := strings.NewReader(string(json))
	config, err := NewConfig(reader)
	ok(t, err)

	return expected, config
}

func TestConfigEmail(t *testing.T) {
	expected, config := createConfig(t)
	equals(t, expected.Email, config.Email)
}

func TestConfigPassword(t *testing.T) {
	expected, config := createConfig(t)
	equals(t, expected.Password, config.Password)
}
