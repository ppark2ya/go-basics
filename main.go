package main

import (
	"errors"
	"fmt"
	"net/http"
)

var errRequestFailed = errors.New("Request failed")

func main() {
	// make empty map
	// same: var results = map[string]string{}
	var results = make(map[string]string)

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
		result := "OK"
		err := hitUrl(url)

		if err != nil {
			result = "FAILED"
		}
		results[url] = result
	}

	for url, result := range results {
		fmt.Println(url, result)
	}
}

func hitUrl(url string) error {
	fmt.Println("CHECKING:", url)
	resp, err := http.Get(url)

	if err != nil || resp.StatusCode >= 400 {
		fmt.Println(err, resp.StatusCode)
		return errRequestFailed
	}

	return nil
}
