package main

import (
	"fmt"
	"net/http"
	"strconv"
)

type requestResult struct {
	url string
	status string
}

func main() {
	results := make(map[string]string)
	c := make(chan requestResult)
	urls := []string{
		"https://www.google.com/",
		"https://www.naver.com/",
		"https://www.daum.net/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
	}

	for _, url := range urls {
		go hitUrl(url, c)
	}

	for i := 0; i < len(urls); i++ {
		result := <-c
		results[result.url] = result.status
	}

	for url, status := range results {
		fmt.Println(url, status)
	}
}

// send only channel
func hitUrl(url string, c chan <- requestResult) {
	resp, err := http.Get(url)
	statusCode := strconv.Itoa(resp.StatusCode)

	if err != nil || resp.StatusCode >= 400 {
		c <- requestResult{
			url: url,
			status: statusCode + ": FAILED",
		}
	} else {
		c <- requestResult{
			url: url,
			status: statusCode + ": OK",
		}
	}
}
