package crawl

import (
	"fmt"
	"net/url"
	"time"
)


// Create a type that will implement fetching
// Alongside trackers
type AsHTTPCrawler struct {
	fetcher Fetcher // ./check fetcher.go
	tracker Tracker // ./tracker.go
	workers []workers
	seedUrl *url.URL
}

// NewAsHTTPCrawler takes in a Fetcher
// Fetcher starts the crawls. It also takes in
// [] workers that processes and create a sitemap 
func NewAsHTTPCrawler(seedUrl *url.URL) *AsHTTPCrawler {
	fetcher = NewAsHTTPFetcher()
}

// The crawler is responsible for
// crawling a domain, returning the domains
// that were crawled.
type Crawler interface {
	Crawl(url string) (sitemap.Sitemapper, error)
}

// Implementation of Crawler Interface 
