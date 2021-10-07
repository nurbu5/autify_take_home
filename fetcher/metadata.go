package fetcher

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
)

// PrintMetadata prints metadata associated to the response for the given fetcher
func (f Fetcher) PrintMetadata() {
	f.printURL()
	f.printLastFetchTime()
	f.printLinkCount()
	f.printImageCount()
}

// printURL determines what URL the request was made to and prints that to the console
func (f Fetcher) printURL() {
	u := f.resp.Request.URL
	fmt.Printf("Site: %s\n", u)
}

// printLastFetchTime determines when the response was returned and prints that to the console
func (f Fetcher) printLastFetchTime() {
	t := f.resp.Header["Date"][0]
	fmt.Printf("Last fetch: %s\n", t)
}

// printLinkCount determines how many "a" tags are used in the body of the fetched response
func (f Fetcher) printLinkCount() {
	doc, err := goquery.NewDocumentFromReader(f.resp.Body)
	check(err)
	c := len(doc.Find("a").Nodes)
	fmt.Printf("Link count: %d\n", c)
}

// printImageCount determines how many "img" tags are used in the body of the fetched response
func (f Fetcher) printImageCount() {
	doc, err := goquery.NewDocumentFromReader(f.resp.Body)
	check(err)
	c := len(doc.Find("img").Nodes)
	fmt.Printf("Image count: %d\n", c)
}
