package main

import (
	"io"
	"net/http"
	"sync"
)

type Page struct {
	URL  string
	Body string
}

func Scrape(url string) Page {
	resp , _ := http.Get(url)
	body ,_ := io.ReadAll(resp.Body)
	return Page{URL: url, Body: string(body)}
}

func main() {
	urls := []string{"https://example.com", "https://google.com", "https://github.com"}

	var wg sync.WaitGroup
	result := make(chan Page, len(urls))

	for _, url := range urls {
		wg.Add(1)
		go func(url string){
			defer wg.Done()
			result <- Scrape(url)
		}(url)
	}


}
