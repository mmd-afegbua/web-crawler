package crawl

import (
	"fmt"
	"net/http"
	"net/url"
	"github.com/mmd-afegbua/web-crawler/utils"
)

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
// It does not start the Fetcher; run does
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

// Implementation of the interface
func (a *AsHTTPFetcher) Fetch(url *url.URL) error {
	if a.AsWorker.State() == STOPPED {
		return fmt.Error("%s is in STOPPED state", a.AsWorker.Type())
	}
	if err := a.validate(url); err != nil {
		return err
	}
	normURL, _ := util.NormalizeURL(url)
	util.Printf("Fetcher: Adding URL %v to request queue\n", normURL)
	*a.requestQueue <- *normURL
	return nil
}

func (a *AsHTTPFetcher) ResponseChannel() (responeQueue *FetchResponseQueue) {
	return a.responeQueue
}

func (a *AsHTTPFetcher) Worker() Worker {
	return a.AsWorker
}

// Run starts a loop that sits there and wait for requests
// or the quit signal. It'll be interrupted once 
// the stop method is used
func (a *AsHTTPFetcher) Run() error {
	s.AsWorker.SetState(RUNNING)
	for {
		a.AsWorker.SetState(WAITING)
		select {
		// A request is received
		case req := <-*a.requestQueue:
			a.AsWorker.SetState(RUNNING)
			res, err := a.client.Get(req.String())
			*a.responeQueue <- &FetchMessage {
				Request: 	&req,
				Response: 	res,
				Error:		err,
			}
		// When a quit is received and stop is invoked
		case <-a.AsWorker.Quit:
			a.Worker().SetState(STOPPED)
			return nil
		default:	
		}
	}
}

func (a *AsHTTPFetcher) validate(uri *url.URL) error {
	if uri.Scheme != "http" && uri.Scheme != "https" {
		return fmt.Errorf("Unsupported uri scheme %s", uri.Scheme)
	}
	return nil
}