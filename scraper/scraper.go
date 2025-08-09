package scraper

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"
)

type Fields struct {
	Title       bool
	Channel     bool
	Views       bool
	UploadDate  bool
	Description bool
	Thumbnail   bool
}

type Metadata struct {
	Title       string `json:"title,omitempty"`
	Channel     string `json:"channel,omitempty"`
	Views       string `json:"views,omitempty"`
	UploadDate  string `json:"uploadDate,omitempty"`
	Description string `json:"description,omitempty"`
	Thumbnail   string `json:"thumbnail,omitempty"`
	URL         string `json:"url,omitempty"`
}

func ScrapeMetadata(url string, fields Fields) (*Metadata, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch URL: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read body: %w", err)
	}

	body := string(bodyBytes)
	meta := &Metadata{URL: url}

	// Title
	if fields.Title {
		if title := extract(body, `"title":{"runs":\[{"text":"(.*?)"}`); title != "" {
			meta.Title = htmlUnescape(title)
		}
	}

	// Channel
	if fields.Channel {
		if channel := extract(body, `"ownerChannelName":"(.*?)"`); channel != "" {
			meta.Channel = htmlUnescape(channel)
		} else if channel := extract(body, `"channelName":{"simpleText":"(.*?)"`); channel != "" {
			meta.Channel = htmlUnescape(channel)
		}
	}

	// Views (more robust pattern)
	if fields.Views {
		if views := extract(body, `"viewCount":"(\d+)"`); views != "" {
			meta.Views = views
		} else if views := extract(body, `"viewCountText":{"simpleText":"([\d,]+) views"`); views != "" {
			meta.Views = strings.ReplaceAll(views, ",", "")
		} else if views := extract(body, `"shortViewCountText":{"simpleText":"(.*?) views"`); views != "" {
			meta.Views = views
		}
	}

	// Upload Date
	if fields.UploadDate {
		if date := extract(body, `"dateText":{"simpleText":"(.*?)"`); date != "" {
			meta.UploadDate = date
		}
	}

	// Description
	if fields.Description {
		if desc := extract(body, `"description":{"simpleText":"(.*?)"`); desc != "" {
			meta.Description = htmlUnescape(desc)
		}
	}

	// Thumbnail
	if fields.Thumbnail {
		if thumb := extract(body, `"thumbnail":{"thumbnails":\[{"url":"(.*?)"`); thumb != "" {
			meta.Thumbnail = thumb
		}
	}

	if isAllEmpty(meta) {
		return nil, errors.New("no metadata found")
	}

	return meta, nil
}

// extract finds the first match for a regex pattern in the given text.
func extract(text, pattern string) string {
	re := regexp.MustCompile(pattern)
	match := re.FindStringSubmatch(text)
	if len(match) > 1 {
		return match[1]
	}
	return ""
}

func htmlUnescape(s string) string {
	r := strings.NewReplacer(
		`&amp;`, "&",
		`&lt;`, "<",
		`&gt;`, ">",
		`&quot;`, `"`,
		`&#39;`, "'",
	)
	return r.Replace(s)
}

func isAllEmpty(meta *Metadata) bool {
	return meta.Title == "" &&
		meta.Channel == "" &&
		meta.Views == "" &&
		meta.UploadDate == "" &&
		meta.Description == "" &&
		meta.Thumbnail == ""
}
