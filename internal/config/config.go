package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

const defaultURL = "https://ip-ranges.amazonaws.com/ip-ranges.json"

var Config struct {
	AWSIPRangesURL string `json:"aws_ip_ranges_url"`
}

func LoadConfig() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("❌ Error getting user home directory:", err)
		Config.AWSIPRangesURL = defaultURL
		return
	}

	configPath := filepath.Join(homeDir, ".awsipsrc")
	file, err := os.Open(configPath)
	if err != nil {
		fmt.Println("ℹ️ No config file found, using default AWS IP ranges URL.")
		fmt.Println("💡 Tip: Create a config file in your home directory named .awsipsrc with the key aws_ip_ranges_url in JSON format to customize the source.")
		Config.AWSIPRangesURL = defaultURL
		return
	}
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Println("⚠️ Warning: Failed to close config file:", err)
		}
	}()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&Config); err != nil {
		fmt.Println("❌ Error parsing config file:", err)
		Config.AWSIPRangesURL = defaultURL
	}
}
