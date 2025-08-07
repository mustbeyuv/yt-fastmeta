# yt-fastmeta
```
The fastest way to grab YouTube metadata without API keys.
Stop waiting around for slow scrapers or dealing with YouTube API quotas. Get video titles, views, duration, and more in milliseconds. Built for developers who need YouTube data *now*.
```
## yt-fastmeta?
```
Fast. Built in Go with zero external dependencies. No browser automation, no heavy libraries.
Simple. One command. One import. Works everywhere Go works.
No API keys. Skip the YouTube Data API entirely. No quotas, no registration, no hassle.
Flexible. Use it from command line or import into your Go projects.
```

## Quick Start
### Install

```
go install github.com/mustbeyuv/yt-fastmeta@latest
```

### Use

```bash
# Get basic info
yt-fastmeta --url "https://www.youtube.com/watch?v=dQw4w9WgXcQ"
# Get specific fields only  
yt-fastmeta --url "https://www.youtube.com/watch?v=dQw4w9WgXcQ" --fields "title,views"
# Full JSON output
yt-fastmeta --url "https://www.youtube.com/watch?v=dQw4w9WgXcQ" --json
# Using video ID directly
yt-fastmeta --id "dQw4w9WgXcQ"
```

## As a Go Module
Perfect for building YouTube tools, analytics dashboards, or content management systems.
```bash
go get github.com/mustbeyuv/yt-fastmeta
```
## What You Get
Every call returns clean, structured data:
- **Title** - Full video title
- **Views** - View count (formatted)
- **Duration** - Video length 
- **Channel** - Channel name
- **Upload Date** - When it was published
- **Description** - Video description (truncated)
- **Thumbnail** - High-quality thumbnail URL
## CLI Options
```
--url string      YouTube video URL 
--id string       YouTube video ID (alternative to URL)
--json            Output full metadata as JSON
--fields string   Specific fields only (comma-separated)
```

## Available Fields
Use with `--fields` to get only what you need:
`title`, `views`, `duration`, `channel`, `date`, `description`, `thumbnail`
Example: `--fields "title,views,duration"`

## Build from Source
```
git clone https://github.com/mustbeyuv/yt-fastmeta.git
cd yt-fastmeta
go build -o yt-fastmeta main.go
./yt-fastmeta --url "https://www.youtube.com/watch?v=dQw4w9WgXcQ"
```

## Use Cases
**Content creators** - Quick metadata for video analysis
**Developers** - Embed YouTube data in applications  
**Data analysts** - Batch process video information
**Automation** - Integrate with scripts and workflows

## Performance
Typical response times:
- Single video: 100-300ms
- Batch processing: 50+ videos/second
- Memory usage: <10MB

No rate limits. No API keys. Just fast, reliable metadata extraction.

## License
MIT License - use it anywhere, commercial or personal.

## Contributing

Found a bug? Want a feature? Issues and PRs welcome at [github.com/mustbeyuv/yt-fastmeta](https://github.com/mustbeyuv/yt-fastmeta)

---

**Get started:** `go install github.com/mustbeyuv/yt-fastmeta@latest`
