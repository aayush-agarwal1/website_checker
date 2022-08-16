package utils

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func ReadConfig(configFile string) map[string]string {

	properties := make(map[string]string)

	file, err := os.Open(configFile)
	if err != nil {
		log.Printf("Failed to open file: %s", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if !strings.HasPrefix(line, "#") && len(line) != 0 {
			kv := strings.Split(line, "=")
			parameter := strings.TrimSpace(kv[0])
			value := strings.TrimSpace(kv[1])
			properties[parameter] = value
		}
	}

	if err := scanner.Err(); err != nil {
		log.Printf("Failed to read file: %s", err)
		os.Exit(1)
	}

	return properties
}
