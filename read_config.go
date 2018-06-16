package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

func readConfig() (string, error) {
	homePath := os.Getenv("HOME")
	configPath := filepath.Join(homePath, ".config", "sconmpd", "config")
	b, err := ioutil.ReadFile(configPath)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
