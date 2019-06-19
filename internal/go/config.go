package openapi

import (
	"encoding/json"
	"io"
	"os"
)

// Config stores the OurGroceries email and password
type Config struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

// LoadConfig creates a Config from a file
func LoadConfig() (Config, error) {
	file, err := os.Open("config.json")
	if err != nil {
		return Config{}, err
	}
	defer file.Close()
	return NewConfig(file)
}

// NewConfig creates a Config from a stream
func NewConfig(stream io.Reader) (Config, error) {
	decoder := json.NewDecoder(stream)
	config := Config{}
	err := decoder.Decode(&config)
	return config, err
}
