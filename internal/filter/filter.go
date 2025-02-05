package filter

import (
	"awsips/internal/parser"
	"fmt"
	"sort"
)

// FilterByRegion filters IP ranges by the given region and sorts results alphabetically.
func FilterByRegion(ipRanges *parser.AWSIPRanges, region string) {
	var results []string
	found := false

	for _, prefix := range ipRanges.Prefixes {
		if prefix.Region == region {
			results = append(results, fmt.Sprintf("%s: %s", prefix.Service, prefix.IPPrefix))
			found = true
		}
	}

	if !found {
		fmt.Printf("⚠️ No IPs found for region: %s\n", region)
		return
	}

	sort.Strings(results)

	for _, entry := range results {
		fmt.Println(entry)
	}
}

// FilterByService filters IP ranges by the given service and sorts results alphabetically.
func FilterByService(ipRanges *parser.AWSIPRanges, service string) {
	var results []string

	for _, prefix := range ipRanges.Prefixes {
		if prefix.Service == service {
			results = append(results, prefix.IPPrefix)
		}
	}

	sort.Strings(results)

	for _, ip := range results {
		fmt.Println(ip)
	}
}

// FilterByRegionAndService filters IP ranges by both region and service and sorts results alphabetically.
func FilterByRegionAndService(ipRanges *parser.AWSIPRanges, region, service string) {
	var results []string

	for _, prefix := range ipRanges.Prefixes {
		if prefix.Region == region && prefix.Service == service {
			results = append(results, prefix.IPPrefix)
		}
	}

	sort.Strings(results)

	for _, ip := range results {
		fmt.Println(ip)
	}
}

// ListRegions lists all unique regions in the dataset and sorts them alphabetically.
func ListRegions(ipRanges *parser.AWSIPRanges) {
	regionSet := make(map[string]struct{})
	for _, prefix := range ipRanges.Prefixes {
		regionSet[prefix.Region] = struct{}{}
	}

	var regions []string
	for region := range regionSet {
		regions = append(regions, region)
	}

	sort.Strings(regions)

	fmt.Println("🌎 Available Regions:")
	for _, region := range regions {
		fmt.Println(region)
	}
}

// ListServices lists all unique services in the dataset and sorts them alphabetically.
func ListServices(ipRanges *parser.AWSIPRanges) {
	serviceSet := make(map[string]struct{})
	for _, prefix := range ipRanges.Prefixes {
		serviceSet[prefix.Service] = struct{}{}
	}

	var services []string
	for service := range serviceSet {
		services = append(services, service)
	}

	sort.Strings(services)

	fmt.Println("🛠 Available Services:")
	for _, service := range services {
		fmt.Println(service)
	}
}

// ListServicesByRegion lists services available in a specific region.
func ListServicesByRegion(ipRanges *parser.AWSIPRanges, region string) {
	serviceSet := make(map[string]struct{})
	for _, prefix := range ipRanges.Prefixes {
		if prefix.Region == region {
			serviceSet[prefix.Service] = struct{}{}
		}
	}

	if len(serviceSet) == 0 {
		fmt.Printf("⚠️ No services found for region: %s\n", region)
		return
	}

	var services []string
	for service := range serviceSet {
		services = append(services, service)
	}

	sort.Strings(services)

	fmt.Printf("🛠 Services available in %s:\n", region)
	for _, service := range services {
		fmt.Println(service)
	}
}

// ListRegionsByService lists regions where a specific service is available.
func ListRegionsByService(ipRanges *parser.AWSIPRanges, service string) {
	regionSet := make(map[string]struct{})
	for _, prefix := range ipRanges.Prefixes {
		if prefix.Service == service {
			regionSet[prefix.Region] = struct{}{}
		}
	}

	if len(regionSet) == 0 {
		fmt.Printf("⚠️ No regions found for service: %s\n", service)
		return
	}

	var regions []string
	for region := range regionSet {
		regions = append(regions, region)
	}

	sort.Strings(regions)

	fmt.Printf("🌍 Regions where %s is available:\n", service)
	for _, region := range regions {
		fmt.Println(region)
	}
}
