package main

import (
	"flag"
	"fmt"
	"net/http"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var opt string

	urls := []string{
		"http://github.com/lucaslollobrigida",
		"http://google.com",
		"http://spotify.com",
	}

	flag.StringVar(&opt, "m", "", "Execution Mode")
	flag.Parse()

	if opt == "" || opt == "c" {
		concurrent(urls, &wg)
	} else {
		sync_code(urls)
	}
}

func concurrent(urls []string, wg *sync.WaitGroup) {
	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()

			resp, err := http.Get(url)
			if err != nil {
				fmt.Print(err)
				return
			}
			fmt.Println(resp.StatusCode)
		}(url)
	}
	wg.Wait()
}

func sync_code(urls []string) {
	for _, url := range urls {
		func(url string) {
			resp, err := http.Get(url)
			if err != nil {
				fmt.Print(err)
				return
			}
			fmt.Println(resp.StatusCode)
		}(url)
	}
}
