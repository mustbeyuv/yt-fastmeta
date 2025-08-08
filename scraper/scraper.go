package scraper

import (
	"fmt"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

// Metadata holds video metadata
type Metadata struct {
	Title       string
	Channel     string
	Views       string
	UploadDate  string
	Description string
	Thumbnail   string
	URL         string
}

// Fields specifies which metadata fields the user wants
type Fields struct {
	Title       bool
	Channel     bool
	Views       bool
	UploadDate  bool
	Description bool
	Thumbnail   bool
}

// ScrapeMetadata fetches YouTube metadata selectively based on requested fields
func ScrapeMetadata(videoURL string, fields Fields) (*Metadata, error) {
	resp, err := http.Get(videoURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("bad status code: %d", resp.StatusCode)
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, err
	}

	meta := &Metadata{URL: videoURL}

	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode {
			attrs := getAttrs(n)

			// Extract from <meta> tags
			if n.Data == "meta" {
				if fields.Title && attrs["property"] == "og:title" && meta.Title == "" {
					meta.Title = attrs["content"]
				}
				if fields.Thumbnail && attrs["property"] == "og:image" && meta.Thumbnail == "" {
					meta.Thumbnail = attrs["content"]
				}
				if fields.Description && attrs["property"] == "og:description" && meta.Description == "" {
					meta.Description = attrs["content"]
				}
				if fields.UploadDate && attrs["itemprop"] == "uploadDate" && meta.UploadDate == "" {
					meta.UploadDate = attrs["content"]
				}
				if fields.Views && attrs["itemprop"] == "interactionCount" && meta.Views == "" {
					meta.Views = attrs["content"]
				}
				if fields.Channel && attrs["itemprop"] == "name" && meta.Channel == "" {
					meta.Channel = attrs["content"]
				}
			}

			// Also check <link> tag for channel (some pages do this)
			if fields.Channel && n.Data == "link" {
				if attrs["itemprop"] == "name" && meta.Channel == "" {
					meta.Channel = attrs["content"]
				}
			}
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)

	return meta, nil
}

// getAttrs converts node attributes to map
func getAttrs(n *html.Node) map[string]string {
	attrs := make(map[string]string)
	for _, a := range n.Attr {
		attrs[strings.ToLower(a.Key)] = a.Val
	}
	return attrs
}
