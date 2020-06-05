package crawl

import (
	"fmt"
	"net/http"
	"net/url"
	"github.com/mmd-afegbua/web-crawler/utils"
)
// set channel buffer size to 100
const defaultChannelSize = 100

type FetchMessage struct {
	Request *url.URL // For tracking original response
	Response *http.Response // The response
	Error error

}

// Channel thataa listens to incoming requests to the Fetcher
type RequestQueue chan url.URL

// Channel that gets incoming requests
type FetchResponseQueue chan *FetchMessage

type AsHTTPFetcher struct (
	*AsWorker // found in ./worker.go

	requestQueue *RequestQueue
	responeQueue *FetchResponseQueue

	client HTTPClient
)

// NewAsHTTPFetcher is a contructor for AsHTTPFetcher
// It does not start the Fetcher
func NewAsHTTPFetcher() *AsHTTPFetcher {
	reqQueue := make(RequestQueue, defaultChannelSize)
	resQueue := make(FetchResponseQueue, defaultChannelSize)
	a := &AsHTTPFetcher{
		AsWorker: NewAsWorker("Fetcher"),
		client: &http.Client{},
		requestQueue: &reqQueue,
		responeQueue: &resQueue, 
	}
	a.AsWorker.RunFunc = a.Run

	return a
}

// Fetcher is an assynchronous worker interface
// It fetches URLs and exposes the results of type
// FetchMessage through channel ResponseChannel to users
type Fetcher interface {
	Fetch(url *url.URL) error //feeds Fetcher
	
	//ResponseChannel returns the Channel for the Fetcher
	// meant for users
	ResponseChannel() (responeQueue *FetchResponseQueue)
	// Worker manages Fetcher Service
	Worker() Worker
}

// Implementation of the interface