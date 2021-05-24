package main

import (
	"flag"
	"fmt"
	"go_search/pkg/crawler"
	"go_search/pkg/crawler/spider"
	"strings"
)

func main() {
	word := flag.String("s", "", "search by word")
	flag.Parse()

	const depth = 2
	var urls = []string{"https://go.dev", "https://golang.org"}
	docs := scan(urls, depth)
	fmt.Println("Resource scan result: \n", docs)

	if *word == "" { return	}

	fmt.Printf("Search word '%s' in the results:\n", *word)
	find(*word, docs)
}

func scan(urls []string, depth int) (data []crawler.Document) {
	spider := spider.New()

	for _, url := range urls {
		fmt.Println("Scan: ", url)
		docs, err := spider.Scan(url, depth)
		if err != nil {
			fmt.Println("Error: ", err)
			continue
		}
		data = append(data, docs...)
	}

	return data
}

func find(word string, docs []crawler.Document) {
	for _, doc := range docs {
		if strings.Contains(strings.ToLower(doc.Title), word) || strings.Contains(strings.ToLower(doc.URL), word) {
			fmt.Println(doc.URL)
		}
	}
}
