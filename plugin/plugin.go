package plugin

import (
	"log"
	"os"
	"path/filepath"
	"plugin"
)

// Plugin wraps around the plugin.Plugin type also storing the file name
type Plugin struct {
	FileName string
	Plugin   *plugin.Plugin
}

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
			return nil
		}

		plugins = append(plugins, &Plugin{
			FileName: filepath.Base(path),
			Plugin:   p,
		})

		return nil
	})
	if err != nil {
		log.Fatalln("[ERROR]", err.Error())
	}

	return
}
