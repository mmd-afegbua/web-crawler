package crawl

import (
	"fmt"
	"net/url"
	"time"
	"github.com/antoniou/go-crawler/sitemap"
	"github.com/antoniou/go-crawler/util"
)


// NewAsHTTPCrawler takes in a Fetcher
// Fetcher starts the crawls. It also takes in
// [] workers that processes and create a sitemap 
func NewAsHTTPCrawler(seedUrl *url.URL) *AsHTTPCrawler {

	fetcher := NewAsHTTPFetcher()
	parser := NewAsHTTPParser(seedUrl, fetcher)
	tracker := NewAsHTTPTracker(fetcher, parser)
	return &AsHTTPCrawler {
		seedUrl: seedUrl,
		fetcher: fetcher,
		tracker: tracker,
		workers: []Worker{
			parser.Worker(),
			fetcher.Worker(),
			tracker.Worker(),
		},
	}
}

// The crawler is responsible for
// crawling a domain, returning the domains
// that were crawled.
type Crawler interface {
	Crawl(url string) (sitemap.Sitemapper, error)
}

// Create a type that will implement fetching
// Alongside trackers
type AsHTTPCrawler struct {
	fetcher Fetcher // ./fetcher.go
	tracker Tracker // ./tracker.go
	workers []workers
	seedUrl *url.URL
}

// Implementation of Crawler Interface 
func (c *AsHTTPCrawler) Crawl() (sitemap.Sitemapper, error) {
	// Create an empy sitemap
	stmp := sitemap.NewGraphSitemap()
	//Pass the map to the tracker
	c.tracker.SetSitemapper(stmp)

	for _, worker := range c.workers {
		util.Printf("Starting worker of type %v\n", worker.Type())
		go worker.Run()
	}

	fmt.Printf("Starting crawling of %v\n", c.seedURL)
	err := c.fetcher.Fetch(c.seedURL)
	if err != nil {
		return nil, err
	}

	return stmp, c.join()
}

// Wait for all workers to be in state WAITING. This
// will indicate that work is done
func (c *AsyncHTTPCrawler) join() error {
	for {
		time.Sleep(500 * time.Millisecond)
		state := WAITING

		for _, worker := range c.workers {
			state += worker.State()
		}

		if state == WAITING {
			return nil
		}

	}
}