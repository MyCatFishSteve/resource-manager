package provider

import (
	"errors"
	"plugin"
)

// Method symbols to find in a provider
const (
	nameMethod      = "Name"
	versionMethod   = "Version"
	startMethod     = "Start"
	stopMethod      = "Stop"
	terminateMethod = "Terminate"
	setKeyMethod    = "SetKey"
	getKeyMethod    = "GetKey"
)

// Common errors that can occur
var (
	errMissingNameMethod      = errors.New("Missing Name method")
	errMissingVersionMethod   = errors.New("Missing Version method")
	errMissingStartMethod     = errors.New("Missing Start method")
	errMissingStopMethod      = errors.New("Missing Stop method")
	errMissingTerminateMethod = errors.New("Missing Terminate method")
	errMissingSetKeyMethod    = errors.New("Missing SetKey method")
	errMissingGetKeyMethod    = errors.New("Missing GetKey method")
)

// Provider interface
// Name() returns the name of the provider
// Version() returns the
type Provider struct {
	Name    func() string
	Version func() string

	Start     func() error
	Stop      func() error
	Terminate func() error

	SetKey func(string, string)
	GetKey func(string) string
}

// LoadProviders returns a slice of Provider pointers
func LoadProviders(plugins []*plugin.Plugin) (providers []*Provider) {
	for _, plugin := range plugins {
		provider, err := LoadProvider(plugin)
		if err != nil {
			continue
		}

		providers = append(providers, provider)
	}
	return
}

// LoadProvider loads a provider from a plugin
func LoadProvider(plugin *plugin.Plugin) (*Provider, error) {
	name, err := plugin.Lookup(nameMethod)
	if err != nil {
		return nil, errMissingNameMethod
	}

	version, err := plugin.Lookup(versionMethod)
	if err != nil {
		return nil, errMissingVersionMethod
	}

	start, err := plugin.Lookup(startMethod)
	if err != nil {
		return nil, errMissingStartMethod
	}

	stop, err := plugin.Lookup(stopMethod)
	if err != nil {
		return nil, errMissingStopMethod
	}

	terminate, err := plugin.Lookup(terminateMethod)
	if err != nil {
		return nil, errMissingTerminateMethod
	}

	setKey, err := plugin.Lookup(setKeyMethod)
	if err != nil {
		return nil, errMissingSetKeyMethod
	}

	getKey, err := plugin.Lookup(getKeyMethod)
	if err != nil {
		return nil, errMissingGetKeyMethod
	}

	return &Provider{
		Name:      name.(func() string),
		Version:   version.(func() string),
		Start:     start.(func() error),
		Stop:      stop.(func() error),
		Terminate: terminate.(func() error),
		SetKey:    setKey.(func(string, string)),
		GetKey:    getKey.(func(string) string),
	}, nil
}
