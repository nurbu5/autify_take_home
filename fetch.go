package main

import (
	"github.com/nurbu5/fetch/fetcher"
	"log"
	"os"
)

// separateFlags takes a slice of arguments and a map whose keys are made up of valid flag names. The function removes
// any flags used, checks if the flag is valid, updates the given map, and returns any arguments given
func separateFlags(args []string, validFlags map[string]bool) []string {
	urls := make([]string, 0)
	for _, a := range args {
		if len(a) >= 3 && a[:2] == "--" {
			if _, ok := validFlags[a[2:]]; ok != true {
				log.Fatalf("[Error] Invalid flag used: %s\n", a)
			} else {
				validFlags[a[2:]] = true
			}
		} else {
			urls = append(urls, a)
		}
	}
	return urls
}

func main() {
	validFlags := map[string]bool{"metadata": false}
	args := os.Args[1:]
	urls := separateFlags(args, validFlags)
	for _, u := range urls {
		f := fetcher.New(u)
		if validFlags["metadata"] {
			f.PrintMetadata()
		} else {
			f.SaveBody()
		}
	}
}
