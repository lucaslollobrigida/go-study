package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"net/http"
)

func main() {
	var topic string
	flag.StringVar(&topic, "topic", "golang", "Topic to search")
	flag.Parse()

	items, err := Get(topic)
	if err != nil {
		fmt.Print(err)
	}

	for _, item := range items {
		// fmt.Printf("%s -> %s\n", item.Title, item.Url)
		fmt.Println(item.Title)
	}
}

func Get(reddit string) ([]Item, error) {
	url := fmt.Sprintf("https://www.reddit.com/r/%s.json", reddit)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-agent", "golang:reddit_app:v1")
	client := &http.Client{}
	data, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer data.Body.Close()

	if data.StatusCode != http.StatusOK {
		return nil, errors.New(data.Status)
	}

	r := new(apiResponse)

	if err := json.NewDecoder(data.Body).Decode(r); err != nil {
		return nil, err
	}

	items := make([]Item, len(r.Data.Children))

	for i, child := range r.Data.Children {
		items[i] = child.Data
	}

	return items, nil
}

type apiResponse struct {
	Data struct {
		Children []struct {
			Data Item
		}
	}
}

type Item struct {
	Title string
	Url   string
}
