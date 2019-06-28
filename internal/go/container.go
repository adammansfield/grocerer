package openapi

import (
	"net/http"
	"net/http/cookiejar"
)

// Container provides dependencies
type Container struct {
	CookieJar  *cookiejar.Jar
	HTTPClient http.Client
}

// NewContainer creates a Container with real interfaces
func NewContainer() Container {
	cookieJar, _ := cookiejar.New(nil)
	httpClient := http.Client{Jar: cookieJar}
	return Container{cookieJar, httpClient}
}

// TODO: remove `var container` when openapi-generator is no longer used
// The generated code creates functions with parameters that cannot change.
// Since the parameters cannot change dependencies cannot be passed in.
// So use this global variable to inject dependencies instead.
var container = NewContainer()
