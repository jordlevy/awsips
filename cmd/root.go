package cmd

import (
	"awsips/internal/config"
	"awsips/internal/filter"
	"awsips/internal/parser"
	"flag"
	"fmt"
	"log"
	"os"
	"text/tabwriter"
)

func Execute() {
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

	config.LoadConfig()
	ipRanges, err := parser.FetchIPRanges()
	if err != nil {
		log.Fatalf("❌ Error: %v", err)
	}

	switch {
	case *list && *region != "":
		filter.ListServicesByRegion(ipRanges, *region)
	case *list && *service != "":
		filter.ListRegionsByService(ipRanges, *service)
	case *list:
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
	commands := []struct {
		Command     string
		Description string
	}{
		{"awsips --list | -l", "List all available regions and services"},
		{"awsips --region REGION | -r REGION", "Show IPs for a given region"},
		{"awsips --service SERVICE | -s SERVICE", "Show IPs for a given service"},
		{"awsips --region R --service S | -r R -s S", "Show IPs for a specific region and service"},
		{"awsips --help | -h", "Show this help menu"},
	}

	fmt.Println("☁️ awsips - A neat little tool for parsing AWS IPs. ☁️")
	fmt.Println("📌 Usage:")

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	for _, entry := range commands {
		if _, err := fmt.Fprintf(w, "  🔹 %s\t%s\n", entry.Command, entry.Description); err != nil {
			log.Printf("Error writing help output: %v", err)
			return
		}
	}

	if err := w.Flush(); err != nil {
		log.Printf("Error flushing help output: %v", err)
		return
	}
	fmt.Println("\n✨ Easily copy and use AWS IP ranges for your networking needs! ✨")
}
