package main

import (
	"fmt"
	"time"
)

func main() {
	date := time.Now()
	fmt.Printf("UTC: %v\nUnix: %v\nISO: %v\n", date.UTC(), date.Unix(), date.UTC().Format("2006-01-02T15:04:05-0700"))
}
