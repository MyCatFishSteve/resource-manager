package main

import (
	"log"
	"os"
	"path/filepath"
)

// Returns the configuration directory
func configDir() string {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatalln("[ERROR] Failed to get user directory:", err.Error())
	}

	return filepath.Join(home, ".resource-manager.d")
}

// Return the plugin directory
func pluginDir() string {
	return filepath.Join(configDir(), "plugins")
}

func createDirs() error {
	if err := os.Mkdir(configDir(), os.ModePerm); err != nil {
		return err
	}

	if err := os.Mkdir(pluginDir(), os.ModePerm); err != nil {
		return err
	}

	return nil
}
