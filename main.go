package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/mustbeyuv/yt-fastmeta/internal/scraper"
	"github.com/mustbeyuv/yt-fastmeta/internal/search"
)

const version = "v0.1.0" 

func main() {
	urlFlag := flag.String("url", "", "YouTube video URL to fetch metadata for")
	searchFlag := flag.String("search", "", "Search query to fetch video URLs")
	limit := flag.Int("limit", 1, "Number of search results to return (used with --search)")
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
		outputJSON(results)
		return

	case *urlFlag != "":
		meta, err := scraper.ScrapeMetadata(*urlFlag, scraper.Fields{
			Title:       true,
			Channel:     true,
			Views:       true,
			UploadDate:  true,
			Description: true,
			Thumbnail:   true,
		})
		if err != nil {
			log.Fatalf("scrape error: %v", err)
		}
		outputJSON(meta)
		return

	default:
		printHelp()
	}
}

func printHelp() {
	fmt.Println(`yt-fastmeta - A fast YouTube metadata fetcher

Usage:
  yt-fastmeta --url "<youtube_video_url>"
  yt-fastmeta --search "lofi chill" --limit 3

Flags:
  --url       YouTube video URL
  --search    YouTube search query
  --limit     (optional) Number of results to fetch (default 1)
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
