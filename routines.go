package main

import (
	"fmt"
	//"io/ioutil"
	"net/http"
	"sync"
)

func processUrl(url string) {
	response, err := http.Get(url)

	defer response.Body.Close()
	if err == nil {
		fmt.Println(response.Status)
	}
	wg.Done()
}

func main() {

	urls := []string{"http://www.google.com", "http://www.yahoo.com", "http://www.google.com", "http://www.yahoo.com", "http://www.google.com", "http://www.yahoo.com"}

	for _, v := range urls {
		go processUrl(v)
		wg.Add(1)
	}

	wg.Wait()
}
