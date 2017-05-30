package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, visited Visited) {
	// This implementation doesn't do either:
	if depth <= 0 {
		return
	}
	if !visited.insertIfNotPresent(url) {
		// url not inserted -> url already visited
		return
	}
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		wg.Add(1)
		// TODO - Without passing u as an argument 'cause unlike Java, go will pick up whatever value of
		// TODO - "u" is when go is invoked
		go func (u string) {
			defer wg.Done()
			Crawl(u, depth-1, fetcher, visited)
		}(u)
	}
	return
}

func main() {
	visited := Visited{visitedNodes: make(map[string]bool)}
	wg.Add(1)
	go func () {
		defer wg.Done()
		Crawl("http://golang.org/", 4, fetcher, visited)
	}()
	wg.Wait()
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type Visited struct {
	visitedNodes map[string]bool
	rwMutex sync.RWMutex
};

func (v *Visited) insertIfNotPresent(url string)  bool {
	if present := v.contains(url); !present {
		v.rwMutex.Lock()
		defer v.rwMutex.Unlock()
		v.visitedNodes[url] = true
		return true // inserted
	}
	return false // not inserted
}

func (v *Visited) contains(url string)  bool {
	v.rwMutex.RLock()
	defer v.rwMutex.RUnlock()
	_, ok := v.visitedNodes[url]
	return ok
}

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
