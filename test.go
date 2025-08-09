import (
	"github.com/mustbeyuv/yt-fastmeta/scraper"
	"github.com/mustbeyuv/yt-fastmeta/search"
)

func Example() {
	results, _ := search.Search("lofi", 3)

	for _, url := range results {
		meta, _ := scraper.ScrapeMetadata(url, scraper.Fields{
			Title:       true,
			Channel:     true,
			Views:       true,
			UploadDate:  true,
			Description: true,
			Thumbnail:   true,
		})
		fmt.Println(meta.Title)
	}
}
