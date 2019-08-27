package main

import (
	"github.com/immediate-media/resource-manager/plugin"
	"github.com/immediate-media/resource-manager/provider"
	"log"
	"os"
)

// Ensure required resources exist
func init() {
	createDirs()
}

func main() {
	log.SetPrefix("")
	os.Exit(realMain())
}

func realMain() int {
	plugins := plugin.LoadPluginDir(pluginDir())
	var providers []*provider.Provider

	for _, plugin := range plugins {
		provider, err := provider.LoadProvider(plugin.Plugin)
		if err != nil {
			log.Fatalln("[ERROR]", err.Error(), "from", plugin.FileName)
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

	switch action {
	case "start":
		startProviders(providers)
	case "stop":
		stopProviders(providers)
	case "terminate":
		terminateProviders(providers)
	default:
		log.Println("Action not recognised")
	}

	return 0
}

func startProviders(providers []*provider.Provider) {
	for idx, provider := range providers {
		log.Printf("(%d/%d) Start signal sent to %s", idx+1, len(providers), provider.Name())
		provider.Start()
	}
}

func stopProviders(providers []*provider.Provider) {
	for idx, provider := range providers {
		log.Printf("(%d/%d) Stop signal sent to %s", idx+1, len(providers), provider.Name())
		provider.Stop()
	}
}

func terminateProviders(providers []*provider.Provider) {
	for idx, provider := range providers {
		log.Printf("(%d/%d) Terminate signal sent to %s", idx+1, len(providers), provider.Name())
		provider.Terminate()
	}
}
