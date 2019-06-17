package openapi

// Container provides dependencies
type Container struct {
	Config Config
}

// NewContainer creates a Container with real interfaces
func NewContainer() Container {
	config, _ := LoadConfig()
	return Container{config}
}

// TODO: remove `var container` when openapi-generator is no longer used
var container = NewContainer()
