package fetcher

import (
	"fmt"
	"strings"
)

// sanitize cleans a url string so that it can be used as a file or directory name
func sanitize(host, path string) string {
	s := fmt.Sprintf("%s%s", host, path[:len(path)-1])
	s = strings.ReplaceAll(s, "/", "[forwardslash]")
	return s
}
