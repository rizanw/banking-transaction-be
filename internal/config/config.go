package config

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

func New(appName string) (*Config, error) {
	fileConfig := getConfigFile(appName)

	f, err := os.Open(fileConfig)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var cfg Config
	if err = yaml.NewDecoder(f).Decode(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func getConfigFile(appName string) string {
	var (
		filename = fmt.Sprintf("%s/config.yaml", appName)
	)

	dir, _ := os.Getwd()

	return filepath.Join(dir, "files/etc", filename)
}
