package main

import (
	"errors"
	"fmt"
	"sync"
)

type Fetcher interface {
	Fetch(url string) (body string, urls []string, err error)
}

var (
	loading    = errors.New("url is loading")
	urlHistory = struct {
		sync.Mutex
		m map[string]error
	}{m: make(map[string]error)}
)

func Crawl(url string, depth int, fetcher Fetcher) {
	if depth <= 0 {
		return
	}
	urlHistory.Lock()
	_, ok := urlHistory.m[url]
	if ok {
		urlHistory.Unlock()
		return
	}
	urlHistory.m[url] = loading
	urlHistory.Unlock()
	body, urls, err := fetcher.Fetch(url)
	urlHistory.Lock()
	urlHistory.m[url] = err
	urlHistory.Unlock()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	done := make(chan bool)
	for _, u := range urls {
		go func(url string) {
			Crawl(url, depth-1, fetcher)
			done <- true
		}(u)
	}
	for range urls {
		<-done
	}
	return
}

func main() {
	Crawl("https://golang.org/", 4, fetcher)
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
