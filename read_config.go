package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func readConfig() (string, error) {
	homePath := os.Getenv("HOME")
	configPath := filepath.Join(homePath, ".config", "sconmpd", "config")
	b, err := ioutil.ReadFile(configPath)
	if err != nil {
		return ":6600", err
	}
	n := strings.Split(string(b), "\n")
	return n[0], nil
}
