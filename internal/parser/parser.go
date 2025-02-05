package parser

import (
	"awsips/internal/config"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
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
	client := &http.Client{Timeout: 10 * time.Second} // Timeout for reliability

	fmt.Println("📡 Fetching AWS IP ranges from", config.Config.AWSIPRangesURL)
	resp, err := client.Get(config.Config.AWSIPRangesURL)
	if err != nil {
		return nil, fmt.Errorf("❌ Failed to fetch AWS IP ranges: %w", err)
	}
	defer resp.Body.Close()

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
