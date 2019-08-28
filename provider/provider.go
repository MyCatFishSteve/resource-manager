package provider

import (
	"errors"
	"plugin"
)

// Method symbols to find in a provider
const (
	nameMethod    = "Name"
	versionMethod = "Version"
	actionMethod  = "Action"
	setKeyMethod  = "SetKey"
	getKeyMethod  = "GetKey"
)

// Common errors that can occur
var (
	errInvalidProvider = errors.New("Invalid provider interface")
)

// Provider interface
type Provider interface {
	Name() string          // Returns the name of the provider
	Version() string       // Returns the version of the provider
	Action(string) error   // Start an action on the provider
	SetKey(string, string) // Set a key on the plugin
	GetKey(string) string  // Get a key from the plugin
}

// LoadProvider returns a slice of Provider pointers
func LoadProvider(plugin *plugin.Plugin) (Provider, error) {
	symProvider, err := plugin.Lookup("Provider")
	if err != nil {
		return nil, err
	}

	var provider Provider
	provider, ok := symProvider.(Provider)
	if !ok {
		return nil, errInvalidProvider
	}

	return provider, nil
}
