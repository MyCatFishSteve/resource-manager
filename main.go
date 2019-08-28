package main

import (
	"log"
	"os"

	"github.com/immediate-media/resource-manager/plugin"
	"github.com/immediate-media/resource-manager/provider"
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

	// Iterate through the slice of loaded plugins and load them
	// as providers
	for _, plugin := range plugins {
		provider, err := provider.LoadProvider(plugin.Plugin)
		if err != nil {
			log.Fatalln("[ERROR]", err.Error(), "from", plugin.Name)
		}

		providers = append(providers, provider)
	}

	for _, p := range providers {
		log.Println("Name:", p.Name())
		log.Println("Version:", p.Version())
	}

	if len(os.Args) < 2 {
		log.Fatalln("No action was provided")
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
