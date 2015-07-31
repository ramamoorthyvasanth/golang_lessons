package main

import (
	"flag"
	"fmt"
	"github.com/jackdanger/collectlinks"
	//"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

var visited = make(map[string]bool)

func main() {
	flag.Parse()
	args := flag.Args()

	if len(args) < 1 {
		fmt.Println("Please input the URL")
		os.Exit(1)
	}
	queue := make(chan string)

	go func() { queue <- args[0] }()

	for uri := range queue {
		enqueue(uri, queue)
	}

}

func enqueue(url string, q chan string) {
	fmt.Println("fetching", url)
	res, err := http.Get(url)
	if err != nil {
		return
	}
	defer res.Body.Close()

	visited[url] = true
	links := collectlinks.All(res.Body)
	for _, link := range links {
		aURL := getAbsoluteURL(link, url)
		if !visited[aURL] {
			go func() { q <- aURL }()
		}
	}
}

func getAbsoluteURL(href, base string) string {
	url, err := url.Parse(href)
	if err != nil {
		return ""
	}
	baseurl, err := url.Parse(base)
	if err != nil {
		return ""
	}
	uri := baseurl.ResolveReference(url)
	return uri.String()
}
