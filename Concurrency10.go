package main

import (
	"fmt"
	"sync"
)

type UrlCache struct {
	mutex sync.Mutex
	cache map[string]string
}

func (c *UrlCache) Add(url, body string) {
	c.mutex.Lock()
	c.cache[url] = body
	c.mutex.Unlock()
}

func (c *UrlCache) Get(url string) (string, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	v, e := c.cache[url]
	return v, e
}

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, cache UrlCache, ch chan string) {
	// TODO: Fetch URLs in parallel.
	// TODO: Don't fetch the same URL twice.
	// This implementation doesn't do either:
	if depth <= 0 {
		return
	}
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	ch <- url
	cache.Add(url, body)
	
	for _, u := range urls {
		if _, exist := cache.Get(u); !exist {
			go Crawl(u, depth-1, fetcher, cache, ch)
		}
	}
	return
}

const DEPTH = 4

func main() {
	urlChannel := make(chan string)
	cache := UrlCache{ cache: make(map[string]string) }
	go Crawl("http://golang.org/", DEPTH, fetcher, cache, urlChannel)
	for i := 0; i < DEPTH; i++ {
		fmt.Printf("found: %s\n", <- urlChannel)
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
	"http://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"http://golang.org/pkg/",
			"http://golang.org/cmd/",
		},
	},
	"http://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"http://golang.org/",
			"http://golang.org/cmd/",
			"http://golang.org/pkg/fmt/",
			"http://golang.org/pkg/os/",
		},
	},
	"http://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
	"http://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
}
