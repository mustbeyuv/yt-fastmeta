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
	fieldsFlag := flag.String("fields", "", "Comma-separated list of metadata fields to fetch (title,channel,views,uploadDate,description,thumbnail)")
	help := flag.Bool("help", false, "Show usage")
	versionFlag := flag.Bool("version", false, "Print version")
	flag.Parse()

	switch {
	case *help:
		printHelp()
		return

	case *versionFlag:
		fmt.Println("yt-fastmeta version", version)
		return

	case *searchFlag != "":
		results, err := search.Search(*searchFlag, *limit)
		if err != nil {
			log.Fatalf("search error: %v", err)
		}
		if len(results) == 0 {
			log.Println("no results found for the search query")
			os.Exit(1)
		}
		outputJSON(results)
		return

	case *urlFlag != "":
		fields := parseFields(*fieldsFlag)
		// If no fields specified, default to all true
		if fields == (scraper.Fields{}) {
			fields = scraper.Fields{
				Title:       true,
				Channel:     true,
				Views:       true,
				UploadDate:  true,
				Description: true,
				Thumbnail:   true,
			}
		}

		meta, err := scraper.ScrapeMetadata(*urlFlag, fields)
		if err != nil {
			log.Fatalf("scrape error: %v", err)
		}
		if meta.Title == "" && meta.Channel == "" {
			log.Println("no metadata found for the provided URL")
			os.Exit(1)
		}
		outputJSON(meta)
		return

	default:
		printHelp()
	}
}

func parseFields(s string) scraper.Fields {
	f := scraper.Fields{}
	if s == "" {
		return f
	}
	parts := strings.Split(s, ",")
	for _, p := range parts {
		switch strings.ToLower(strings.TrimSpace(p)) {
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

func printHelp() {
	fmt.Println(`yt-fastmeta - A fast YouTube metadata fetcher

Usage:
  yt-fastmeta --url "<youtube_video_url>" [--fields "title,views"]
  yt-fastmeta --search "lofi chill" --limit 3

Flags:
  --url       YouTube video URL
  --search    YouTube search query
  --limit     (optional) Number of results to fetch (default 1)
  --fields    (optional) Comma-separated metadata fields to fetch (title,channel,views,uploadDate,description,thumbnail)
  --version   Print tool version
  --help      Show this help message`)
}

func outputJSON(v interface{}) {
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	if err := enc.Encode(v); err != nil {
		log.Fatalf("json error: %v", err)
	}
}
