package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
)

type Page struct {
	URL  string
	Body string
}

func Scrape(url string) (Page, error) {
	resp, err := http.Get(url)
	if err !=nil{
		return Page{URL: url}, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Page{URL: url}, err
	}
	return Page{URL: url, Body: string(body)}, nil
}

func main() {
	urls := []string{"https://example.com", "https://google.com", "https://github.com"}

	var wg sync.WaitGroup
	results := make(chan Page, len(urls))

	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			page, err := Scrape(url)
			if err != nil {
				fmt.Printf("Error scraping %s: %v\n", url, err)
				return
			}
			results <- page
		}(url)
	}

	go func() { wg.Wait(); close(results) }()

	for p := range results {
		fmt.Printf("%s → %d bytes\n", p.URL, len(p.Body))
	}

}
