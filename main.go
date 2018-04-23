package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func makeRequest(url string) {
	start := time.Now()

	res, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	elapsed := time.Since(start).Seconds()

	fmt.Printf("%s took %v seconds \n", url, elapsed)
}

func webRequests() {
	urls := []string{"https://www.usa.gov/", "https://www.gov.uk", "http://www.gouvernement.fr/"}

	for _, u := range urls {
		// the keyword 'go' makes it run in another thread
		go makeRequest(u)
	}
}

func main() {
	webRequests()

	time.Sleep(time.Second * 5)
}
