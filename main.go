package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/mustbeyuv/yt-fastmeta/scraper"
	"github.com/mustbeyuv/yt-fastmeta/search"
)

const version = "v0.1.0"

func main() {
	urlFlag := flag.String("url", "", "YouTube video URL to fetch metadata for")
	searchFlag := flag.String("search", "", "Search query to fetch video URLs")
	limit := flag.Int("limit", 1, "Number of search results to return (used with --search)")
	fieldsFlag := flag.String("fields", "title,channel,views,uploadDate,description,thumbnail", "Comma-separated list of metadata fields to fetch (title,channel,views,uploadDate,description,thumbnail)")
	jsonFlag := flag.Bool("json", true, "Output in JSON format (default true)")
	help := flag.Bool("help", false, "Show usage")
	versionFlag := flag.Bool("version", false, "Print version")

	flag.Parse()

	if *help {
		printHelp()
		return
	}

	if *versionFlag {
		fmt.Println("yt-fastmeta version", version)
		return
	}

	// Parse requested fields
	fields := parseFields(*fieldsFlag)

	switch {
	case *searchFlag != "":
		results, err := search.Search(*searchFlag, *limit)
		if err != nil {
			log.Fatalf("search error: %v", err)
		}
		if len(results) == 0 {
			log.Println("no results found for the search query")
			os.Exit(1)
		}
		output(results, *jsonFlag)
		return

	case *urlFlag != "":
		meta, err := scraper.ScrapeMetadata(*urlFlag, fields)
		if err != nil {
			log.Fatalf("scrape error: %v", err)
		}
		// Check if we got any meaningful data in requested fields
		if isEmptyMetadata(meta, fields) {
			log.Println("no metadata found for the provided URL")
			os.Exit(1)
		}
		output(meta, *jsonFlag)
		return

	default:
		printHelp()
	}
}

func parseFields(fieldsStr string) scraper.Fields {
	f := scraper.Fields{}
	for _, field := range strings.Split(fieldsStr, ",") {
		switch strings.ToLower(strings.TrimSpace(field)) {
		case "title":
			f.Title = true
		case "channel":
			f.Channel = true
		case "views":
			f.Views = true
		case "uploaddate":
			f.UploadDate = true
		case "description":
			f.Description = true
		case "thumbnail":
			f.Thumbnail = true
		}
	}
	return f
}

func isEmptyMetadata(meta *scraper.Metadata, fields scraper.Fields) bool {
	// Return true if all requested fields are empty strings
	if fields.Title && meta.Title != "" {
		return false
	}
	if fields.Channel && meta.Channel != "" {
		return false
	}
	if fields.Views && meta.Views != "" {
		return false
	}
	if fields.UploadDate && meta.UploadDate != "" {
		return false
	}
	if fields.Description && meta.Description != "" {
		return false
	}
	if fields.Thumbnail && meta.Thumbnail != "" {
		return false
	}
	return true
}

func output(v interface{}, jsonOut bool) {
	if jsonOut {
		enc := json.NewEncoder(os.Stdout)
		enc.SetIndent("", "  ")
		if err := enc.Encode(v); err != nil {
			log.Fatalf("json error: %v", err)
		}
	} else {
		// Plain text fallback for Metadata or []string (search results)
		switch val := v.(type) {
		case *scraper.Metadata:
			fmt.Println("Metadata:")
			if val.Title != "" {
				fmt.Println("Title:", val.Title)
			}
			if val.Channel != "" {
				fmt.Println("Channel:", val.Channel)
			}
			if val.Views != "" {
				fmt.Println("Views:", val.Views)
			}
			if val.UploadDate != "" {
				fmt.Println("Upload Date:", val.UploadDate)
			}
			if val.Description != "" {
				fmt.Println("Description:", val.Description)
			}
			if val.Thumbnail != "" {
				fmt.Println("Thumbnail:", val.Thumbnail)
			}
			fmt.Println("URL:", val.URL)

		case []string:
			fmt.Println("Search Results:")
			for _, url := range val {
				fmt.Println(url)
			}

		default:
			// Unknown type fallback
			fmt.Printf("%v\n", val)
		}
	}
}

func printHelp() {
	fmt.Println(`yt-fastmeta - A fast YouTube metadata fetcher

Usage:
  yt-fastmeta --url "<youtube_video_url>" [--fields "title,views"] [--json=true|false]
  yt-fastmeta --search "lofi chill" --limit 3 [--json=true|false]

Flags:
  --url       YouTube video URL
  --search    YouTube search query
  --limit     Number of results to fetch (default 1)
  --fields    Comma-separated list of metadata fields to fetch (title,channel,views,uploadDate,description,thumbnail)
  --json      Output format: true for JSON (default), false for plain text
  --version   Print tool version
  --help      Show this help message
`)
}
