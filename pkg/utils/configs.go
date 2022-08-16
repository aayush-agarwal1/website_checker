package utils

import (
	"gopkg.in/yaml.v2"
	"log"
	"os"
	"time"
)

type Config struct {
	Server struct {
		Port                string        `yaml:"port"`
		Host                string        `yaml:"host"`
		ReadTimeoutSeconds  time.Duration `yaml:"readTimeoutSeconds"`
		WriteTimeoutSeconds time.Duration `yaml:"writeTimeoutSeconds"`
	} `yaml:"server"`
	Cron struct {
		Concurrency  int           `yaml:"concurrency"`
		DelaySeconds time.Duration `yaml:"delaySeconds"`
	} `yaml:"cron"`
}

func ReadConfig(configFile string) Config {

	var config Config

	file, err := os.Open(configFile)
	if err != nil {
		log.Printf("Failed to open file: %s", err)
		os.Exit(1)
	}
	defer file.Close()

	if err := yaml.NewDecoder(file).Decode(&config); err != nil {
		log.Printf("Incorrect config format %s", err)
		os.Exit(1)
	}
	return config
}
