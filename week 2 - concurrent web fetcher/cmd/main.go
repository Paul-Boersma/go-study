package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	flag.Parse()
	urls := flag.Args()

	results := make()
	fmt.Println("Fetching urls: ", urls)
	for _, url := range urls {
		go fetch(url)
	}

}

func fetch(url string) (time.Duration, error) {
	start := time.Now()

	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	if _, err := io.Copy(io.Discard, resp.Body); err != nil {
		return 0, err
	}

	elapsed := time.Since(start)
	return elapsed, nil
}
