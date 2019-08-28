package plugin

import (
	"log"
	"os"
	"path/filepath"
	"plugin"
)

// Plugin is used to reference a Go plugin and an associated name
type Plugin struct {
	Name   string
	Plugin *plugin.Plugin
}

// NewPlugin returns a Plugin pointer
func NewPlugin(name string, plugin *plugin.Plugin) *Plugin {
	return &Plugin{
		Name:   name,
		Plugin: plugin,
	}
}

// LoadPluginDir loads a directory and returns a slice of Plugins that were found
func LoadPluginDir(root string) (plugins []*Plugin) {
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Println("[ERROR]", err.Error())
			return err
		}

		if info.IsDir() {
			return nil
		}

		p, err := plugin.Open(path)
		if err != nil {
			log.Println("[ERROR]", err.Error())
			return err
		}

		plugins = append(plugins, NewPlugin(filepath.Base(path), p))
		return nil
	})
	if err != nil {
		log.Fatalln("[ERROR]", err.Error())
	}

	return
}
