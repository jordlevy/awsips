package parser

import (
	"awsips/internal/config"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type AWSIPRanges struct {
	Prefixes []struct {
		IPPrefix           string `json:"ip_prefix"`
		Region             string `json:"region"`
		Service            string `json:"service"`
		NetworkBorderGroup string `json:"network_border_group"`
	} `json:"prefixes"`
}

func FetchIPRanges() (*AWSIPRanges, error) {
	fmt.Println("📡 Fetching AWS IP ranges from", config.Config.AWSIPRangesURL)
	resp, err := http.Get(config.Config.AWSIPRangesURL)
	if err != nil {
		return nil, fmt.Errorf("❌ Failed to fetch AWS IP ranges: %w", err)
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Println("⚠️ Warning: Failed to close response body:", err)
		}
	}()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("❌ Received non-200 response: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("❌ Failed to read response body: %w", err)
	}

	var ipRanges AWSIPRanges
	if err := json.Unmarshal(body, &ipRanges); err != nil {
		return nil, fmt.Errorf("❌ Failed to parse JSON: %w", err)
	}

	fmt.Println("✅ Successfully fetched and parsed AWS IP ranges.")
	return &ipRanges, nil
}
