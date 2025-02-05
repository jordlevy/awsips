package cmd

import (
	"awsips/internal/config"
	"awsips/internal/filter"
	"awsips/internal/parser"
	"flag"
	"fmt"
	"log"
)

func Execute() {
	// Define CLI flags with shorthand options
	region := flag.String("region", "", "Filter by AWS region (-r)")
	service := flag.String("service", "", "Filter by AWS service (-s)")
	list := flag.Bool("list", false, "List available regions and services (-l)")
	help := flag.Bool("help", false, "Show help menu (-h)")

	flag.StringVar(region, "r", "", "Shorthand for --region")
	flag.StringVar(service, "s", "", "Shorthand for --service")
	flag.BoolVar(list, "l", false, "Shorthand for --list")
	flag.BoolVar(help, "h", false, "Shorthand for --help")

	flag.Parse()

	if *help || flag.NFlag() == 0 {
		showHelp()
		return
	}

	// Load config and fetch data
	config.LoadConfig()
	ipRanges, err := parser.FetchIPRanges()
	if err != nil {
		log.Fatalf("❌ Error: %v", err)
	}

	// Handle CLI logic using switch
	switch {
	case *list:
		fmt.Println("📜 Listing available regions and services:")
		filter.ListRegions(ipRanges)
		filter.ListServices(ipRanges)
	case *region != "" && *service != "":
		filter.FilterByRegionAndService(ipRanges, *region, *service)
	case *region != "":
		filter.FilterByRegion(ipRanges, *region)
	case *service != "":
		filter.FilterByService(ipRanges, *service)
	default:
		showHelp()
	}
}

func showHelp() {
	fmt.Println(`🛠 awsips - AWS IP Ranges CLI 🛠

📌 Usage:
  🔹 awsips --list | -l                📜 List all available regions and services
  🔹 awsips --region REGION | -r REGION 🌍 Show IPs for a given region
  🔹 awsips --service SERVICE | -s SERVICE 🏗  Show IPs for a given service
  🔹 awsips --region R --service S | -r R -s S 🔍 Show IPs for a specific region and service
  🔹 awsips --help | -h                ℹ️  Show this help menu

✨ Easily copy and use AWS IP ranges for your networking needs! ✨`)
}
