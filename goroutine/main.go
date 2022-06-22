package main

import (
	"log"
	"net/http"
	"time"
)

type result struct {
	Url     string
	Err     error
	Latency time.Duration
}

func get(url string, ch chan<- result) {
	start := time.Now()
	if res, err := http.Get(url); err != nil {
		ch <- result{Url: url, Err: err, Latency: 0}
	} else {
		t := time.Since(start).Round(time.Millisecond)
		ch <- result{Url: url, Err: nil, Latency: t}
		res.Body.Close()
	}
}
func main() {
	results := make(chan result)
	list := []string{"https://google.com", "https://example.com", "https://wiki.sfsf", "https://yahoo.com", "https://google.com/hellworld"}
	for _, url := range list {
		go get(url, results)
	}
	for range list {
		r := <-results
		if r.Err != nil {
			log.Printf("%s \n", r.Err)
		} else {
			log.Printf("%s , %-20d\n", r.Url, r.Latency)
		}
	}
}
