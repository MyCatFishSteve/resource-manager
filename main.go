package main

import (
	"log"
	"os"

	"github.com/immediate-media/resource-manager/plugin"
	"github.com/immediate-media/resource-manager/provider"
)

var (
	supportedActions = []string{
		"start",
		"stop",
		"terminate",
	}
)

// Create required resources
func init() {
	createDirs()
}

func main() {
	log.SetPrefix("")
	os.Exit(realMain())
}

func realMain() int {
	plugins := plugin.LoadPluginDir(pluginDir())
	var providers []provider.Provider

	// Load plugins as providers and add them to providers slice
	for _, plugin := range plugins {
		provider, err := provider.LoadProvider(plugin.Plugin)
		if err != nil {
			log.Fatalln("[ERROR]", err.Error(), "from", plugin.Name)
		}

		providers = append(providers, provider)
	}

	log.Println(len(providers), "loaded into program.")

	if len(os.Args) < 2 {
		log.Println("No action supplied, available actions:")
		for _, action := range supportedActions {
			log.Println("*", action)
		}
	}

	action := os.Args[1]

	actionProviders(action, providers)

	return 0
}

func actionProviders(action string, providers []provider.Provider) {
	for idx, provider := range providers {
		log.Printf("(%d/%d) %s signal sent to %s", idx+1, len(providers), action, provider.Name())
		provider.Action(action)
	}
}
