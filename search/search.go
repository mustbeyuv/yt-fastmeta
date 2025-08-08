package search

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"strings"
)

// Search runs a YouTube search and returns video URLs.
// Example: Search("sad lofi chill", 5)
func Search(query string, limit int) ([]string, error) {
	escaped := url.QueryEscape(strings.TrimSpace(query))
	searchURL := "https://www.youtube.com/results?search_query=" + escaped

	resp, err := http.Get(searchURL)
	if err != nil {
		return nil, fmt.Errorf("http error: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read error: %v", err)
	}
	html := string(body)

	// Extract video IDs from search results
	re := regexp.MustCompile(`/watch\?v=([a-zA-Z0-9_-]{11})`)
	matches := re.FindAllStringSubmatch(html, -1)

	// Deduplicate and limit
	seen := make(map[string]bool)
	results := []string{}

	for _, match := range matches {
		id := match[1]
		if !seen[id] {
			seen[id] = true
			results = append(results, "https://www.youtube.com/watch?v="+id)
			if len(results) >= limit {
				break
			}
		}
	}

	return results, nil
}
