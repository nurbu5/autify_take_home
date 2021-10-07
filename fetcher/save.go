package fetcher

import (
	"io"
	"log"
	"os"
)

// SaveBody creates or overwrites the existing response body file for the
// url associated with the given Fetcher
func (f Fetcher) SaveBody() {
	file, err := os.Create(f.url)
	check(err)
	defer file.Close()

	b := f.processBody()
	_, err = file.Write(b)
	check(err)
}

// processBody reads the response body associated for the given Fetcher
// and reports any errors
func (f Fetcher) processBody() []byte {
	b, err := io.ReadAll(f.resp.Body)
	f.resp.Body.Close()
	if f.resp.StatusCode >= 300 {
		log.Fatalf("Response failed with status code: %d\n and body: %s\n", f.resp.StatusCode, b)
	}
	check(err)
	return b
}
