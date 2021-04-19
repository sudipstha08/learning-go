package main

// provides the ability to embed static files into the binary
// compiled with go build
// Package embed provides access to files embedded in the running Go program.

import (
	_ "embed"
	"log"

	yaml "gopkg.in/yaml.v2"
)

type Config struct {
	Database struct {
		Username string `yaml:"user"`
		Password string `yaml:"pass"`
	} `yaml:"database"`
}

// directive //go:embed will load the configuration file with the given name, 
// and save its content into the variable as a byte array.
//go:embed config.yml
var yamlConfig []byte
var cfg Config

func main() {
	loadConfig()

	// do useful stuff
}

func loadConfig() {
	err := yaml.Unmarshal(yamlConfig, &cfg)
	if err != nil {
		log.Fatal(err)
	}
}


	// https://medium.com/@leo_hetsch/using-gos-embed-package-to-build-a-small-webpage-6175953fccea
