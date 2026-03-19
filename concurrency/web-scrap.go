package main

import (
	"io"
	"net/http"
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

}
