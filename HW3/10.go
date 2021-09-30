package main

import (
	"fmt"
	"sync"
)

var (
	wg = &sync.WaitGroup{}
)

type SafeUrlMap struct {
	fetchedUrls map[string]bool
	mu          *sync.Mutex
}

func (u *SafeUrlMap) SetValue(key string) {
	u.mu.Lock()
	u.fetchedUrls[key] = true
	u.mu.Unlock()
}

func (u *SafeUrlMap) GetValue(key string) (bool, bool) {
	u.mu.Lock()
	defer u.mu.Unlock()
	val, ok := u.fetchedUrls[key]
	return val, ok
}

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func CrawlDfs(url string, depth int, fetcher Fetcher, quit chan bool, safeMap SafeUrlMap) {
	// TODO: Fetch URLs in parallel.
	// TODO: Don't fetch the same URL twice.
	// This implementation doesn't do either:
	defer wg.Done()
	fetched, inMap := safeMap.GetValue(url)
	if fetched && inMap {
		return
	} else {
		safeMap.SetValue(url)
	}

	if depth <= 0 {
		return
	}
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	waitChannel := make(chan bool)
	for _, u := range urls {
		wg.Add(1)
		go CrawlDfs(u, depth-1, fetcher, waitChannel, safeMap)
	}
	return
}

func Crawl() {
	quit := make(chan bool)
	u := SafeUrlMap{fetchedUrls: map[string]bool{}, mu: &sync.Mutex{}}
	wg.Add(1)
	go CrawlDfs("https://golang.org/", 4, fetcher, quit, u)
	wg.Wait()
}

func main() {
	Crawl()
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