package openapi

import (
	"net/http"
	"net/http/cookiejar"
)

// Container provides dependencies
type Container struct {
	CookieJar  *cookiejar.Jar
	Config     Config
	HTTPClient http.Client
}

// NewContainer creates a Container with real interfaces
func NewContainer() Container {
	config, _ := LoadConfig()
	cookieJar, _ := cookiejar.New(nil)
	httpClient := http.Client{Jar: cookieJar}
	return Container{cookieJar, config, httpClient}
}

// TODO: remove `var container` when openapi-generator is no longer used
var container = NewContainer()
