package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

var Data *Config

func Load() error {
	var filePath string
	if os.Getenv("CONFIG_FILE") != "" {
		filePath = os.Getenv("CONFIG_FILE")
	} else {
		filePath = "config.yml"
	}
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}

	defer file.Close()

	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&Data)
	if err != nil {
		return err
	}

	return nil
}
