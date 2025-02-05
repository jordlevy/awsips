package filter

import (
	"awsips/internal/parser"
	"fmt"
)

// FilterByRegion filters IP ranges by the given region.
func FilterByRegion(ipRanges *parser.AWSIPRanges, region string) {
	for _, prefix := range ipRanges.Prefixes {
		if prefix.Region == region {
			fmt.Println(prefix.IPPrefix)
		}
	}
}

// FilterByService filters IP ranges by the given service.
func FilterByService(ipRanges *parser.AWSIPRanges, service string) {
	for _, prefix := range ipRanges.Prefixes {
		if prefix.Service == service {
			fmt.Println(prefix.IPPrefix)
		}
	}
}

// FilterByRegionAndService filters IP ranges by both region and service.
func FilterByRegionAndService(ipRanges *parser.AWSIPRanges, region, service string) {
	for _, prefix := range ipRanges.Prefixes {
		if prefix.Region == region && prefix.Service == service {
			fmt.Println(prefix.IPPrefix)
		}
	}
}

// ListRegions lists all unique regions in the dataset.
func ListRegions(ipRanges *parser.AWSIPRanges) {
	regionSet := make(map[string]struct{})
	for _, prefix := range ipRanges.Prefixes {
		regionSet[prefix.Region] = struct{}{}
	}
	fmt.Println("🌎 Available Regions:")
	for region := range regionSet {
		fmt.Println(region)
	}
}

// ListServices lists all unique services in the dataset.
func ListServices(ipRanges *parser.AWSIPRanges) {
	serviceSet := make(map[string]struct{})
	for _, prefix := range ipRanges.Prefixes {
		serviceSet[prefix.Service] = struct{}{}
	}
	fmt.Println("🛠 Available Services:")
	for service := range serviceSet {
		fmt.Println(service)
	}
}
