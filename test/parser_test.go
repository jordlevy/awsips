package test

import (
	"awsips/internal/parser"
	"encoding/json"
	"strings"
	"testing"
)

func TestParseAWSIPRanges(t *testing.T) {
	var ipRanges parser.AWSIPRanges
	if err := json.NewDecoder(strings.NewReader(SampleJSON)).Decode(&ipRanges); err != nil {
		t.Fatalf("Failed to parse JSON: %v", err)
	}

	if len(ipRanges.Prefixes) == 0 {
		t.Fatal("Parsed IP ranges are empty")
	}
}
