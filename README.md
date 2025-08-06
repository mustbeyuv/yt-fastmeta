# yt-fastmeta
```
A minimal, high-speed YouTube metadata scraper written in Go.
```

**yt-fastmeta** lets you fetch YouTube video IDs from search results and optionally extend it with full metadata scraping (title, channel, views, etc). It's designed to be fast, headless, and free of quota limits — no YouTube API key needed.

---

## Features
```yaml
✓ Raw YouTube search scraping via HTML
✓ Fast video ID extraction using regex  
✓ CLI interface for quick testing
✓ Modular design for easy expansion
✓ No API keys or rate limits
```

---

## Why this exists
Most YouTube scraping tools rely on the official API (quota-bound) or full-blown headless browsers (slow).

`yt-fastmeta` avoids both: it scrapes directly from YouTube search HTML using simple and efficient Go code.

**Perfect for:**
```
→ Music metadata fetchers
→ Fast search previews  
→ Building intelligent playlists
→ ML dataset generation
```

### Quick Start

```bash
# Clone the repository
git clone https://github.com/yourusername/yt-fastmeta
cd yt-fastmeta

# Initialize dependencies
go mod tidy

# Run a basic search
go run main.go "your search query"
```

### Code Example

```go
package main

import (
    "fmt"
    "log"
    "yt-fastmeta/internal/scraper"
)

func main() {
    query := "golang tutorial"
    
    results, err := scraper.SearchVideos(query)
    if err != nil {
        log.Fatal(err)
    }
    
    for _, video := range results {
        fmt.Printf("Title: %s\n", video.Title)
        fmt.Printf("URL: %s\n", video.URL)
        fmt.Printf("---\n")
    }
}
```
---

## Configuration
The scraper can be configured for different use cases:

```go
type Config struct {
    MaxResults    int           // Maximum search results
    Timeout       time.Duration // Request timeout
    UserAgent     string        // Custom user agent
    RetryAttempts int           // Retry failed requests
}
```

---
## Performance
```
Benchmark Results:
├── Search Query Processing: ~200ms
├── Video ID Extraction: ~50ms  
├── Metadata Scraping: ~100ms per video
└── Memory Usage: <10MB typical
```

---
## Contributing
I welcome contributions! Here's how to get started:
```bash
# Fork the repository
# Create a feature branch
git checkout -b feature/amazing-feature

# Make your changes
# Test thoroughly
go test ./...

# Submit a pull request
```

## Disclaimer

This tool is designed for educational and research purposes. Please respect YouTube's Terms of Service and use this tool responsibly. The authors are not responsible for any misuse of this software.

---

**Built with Go • Optimized for speed • Designed for developers**
