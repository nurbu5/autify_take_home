// package fetcher provides a Fetcher struct to fetch data from a url and write
// that data to files locally
package fetcher

import (
	"log"
	"net/http"
)

// Fetcher is a struct that holds a sanitized url string and an HTTP Get response
type Fetcher struct {
	url  string
	resp *http.Response
}

// New generates a new Fetcher with a sanitized url and a http.*Response to a GET request to the given url
func New(url string) Fetcher {
	r := fetchResponse(url)
	return Fetcher{
		url:  sanitize(r.Request.URL.Host, r.Request.URL.Path),
		resp: r,
	}
}

// fetchResponse is a function that runs an HTTP GET on the given url, checks for errors, and returns the response
func fetchResponse(url string) *http.Response {
	r, err := http.Get(url)
	check(err)
	return r
}

// check is a helper function that checks for and logs errors
func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
