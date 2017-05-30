package main

import (
	"fmt"
	"net/http"
	"sync"
)

type response struct {
	res *http.Response;
	err error;
}


func startHttp(req <-chan string, res chan<- response) {
	var wg sync.WaitGroup
	for i := range req {
		wg.Add(1)
		go func(i string, res chan<- response) {
			defer wg.Done()
			resp, err := http.Get(i)
			res <- response{resp, err}
		}(i, res)
	}
	wg.Wait()
	close(res)
}

func main() {
	req := make(chan string, 3)
	res := make(chan response, 3)
	go startHttp(req, res)
	urls := [3] string {"http://www.google.com", "http://www.facebook.com", "http://www.youtube.com/"}

	for i := 0; i< len(urls); i++ {
		fmt.Printf("Fetching url %s\n", urls[i])
		req <- urls[i]
	}
	close(req)

	for i := range res {
		fmt.Printf("*********************************\n")
		fmt.Printf("Response recieved %v\n", i.res)
		fmt.Printf("*********************************\n")
	}
}
