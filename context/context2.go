package main

import (
	"context"
	"net/http"
	"time"
)

type result struct {
	Url     string
	err     error
	latency time.Duration
}

func get(ctx context.Context, url string, ch chan<- result) {
	start := time.Now()
	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if resp, err := http.DefaultClient.Do(req); err != nil {
		ch <- result{Url: url, err: err}
	} else {
		t := time.Since(start).Round(time.Millisecond)
		ch <- result{Url: url, err: nil, latency: t}
		resp.Body.Close()
	}
}
func main() {
	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	results := make(chan result)
	list := []string{
		"https://google.com",
		"https://yahoo.com",
		"https://facebook.com",
		"https://azu.com",
		"https://example.com",
	}
	for _, url := range list {
		go get(ctx, url, results)
	}

}
