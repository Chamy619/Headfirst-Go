package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type Page struct {
	URL string
	Size int
}
func (p Page) String() string {
	return fmt.Sprintf("%s: %d", p.URL, p.Size)
}

func main() {
	sizes := make(chan Page)
	urls := []string{"https://example.com/", "https://go.dev/", "https://go.dev/doc"}
	for _, url := range urls {
		go responseSize(url, sizes)
	}
	for i := 0; i < len(urls); i++ {
		fmt.Println(<-sizes)
	}
}

func responseSize(url string, channel chan Page) {
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	channel <- Page{URL: url, Size: len(body)}
}