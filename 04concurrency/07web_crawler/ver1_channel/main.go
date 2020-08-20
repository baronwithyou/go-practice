package main

import (
	"fmt"
	"sync"
)

// 这里不使用原子map，因为原子性不能保证线程能够取得正确的值
//var fetched sync.Map

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

type fetchState struct {
	fetched map[string]bool
	mu      sync.Mutex
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, result chan<- string, f *fetchState) {
	// TODO: Fetch URLs in parallel.
	// TODO: Don't fetch the same URL twice.
	// This implementation doesn't do either:
	defer close(result)
	if depth <= 0 {
		return
	}

	f.mu.Lock()

	already := f.fetched[url]
	f.fetched[url] = true
	f.mu.Unlock()

	if already {
		return
	}

	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	result <- fmt.Sprintf("[%d] found: %s %q", depth, url, body)

	results := make([]chan string, len(urls))
	for i, u := range urls {
		results[i] = make(chan string)
		go Crawl(u, depth-1, fetcher, results[i], f)
	}

	for _, ret := range results {
		for v := range ret {
			result <- v
		}
	}

	return
}

func main() {
	f := &fetchState{
		fetched: make(map[string]bool),
		mu:      sync.Mutex{},
	}
	//fetched = make(map[string]bool)
	result := make(chan string)
	go Crawl("https://golang.org/", 4, fetcher, result, f)

	for v := range result {
		fmt.Println(v)
	}
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
