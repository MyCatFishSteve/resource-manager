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

	return 0
}
