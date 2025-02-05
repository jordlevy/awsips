package test

import (
	"awsips/internal/filter"
	"awsips/internal/parser"
	"encoding/json"
	"strings"
	"testing"
)

func TestFilterByRegion(t *testing.T) {
	var ipRanges parser.AWSIPRanges
	if err := json.NewDecoder(strings.NewReader(SampleJSON)).Decode(&ipRanges); err != nil {
		t.Fatalf("Failed to parse JSON: %v", err)
	}

	filter.FilterByRegion(&ipRanges, "us-east-1")
}

func TestFilterByService(t *testing.T) {
	var ipRanges parser.AWSIPRanges
	if err := json.NewDecoder(strings.NewReader(SampleJSON)).Decode(&ipRanges); err != nil {
		t.Fatalf("Failed to parse JSON: %v", err)
	}

	filter.FilterByService(&ipRanges, "EC2")
}
