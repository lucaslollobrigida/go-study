package main

import (
	"flag"
	"fmt"
)

type endpoint struct {
	host   string
	port   int
	secure bool
}

func main() {
	e := endpoint{}

	flag.StringVar(&e.host, "host", "", "Host address")
	flag.IntVar(&e.port, "port", 80, "Port")
	flag.BoolVar(&e.secure, "s", false, "HTTPS?")
	flag.Parse()

	if e.secure {
		e.host = fmt.Sprintf("https://%s", e.host)
	} else {
		e.host = fmt.Sprintf("http://%s", e.host)
	}

	fmt.Printf("The call would be like -> %s:%d\n", e.host, e.port)
}
