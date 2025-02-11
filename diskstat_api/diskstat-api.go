package main

import (
	"flag"
	"fmt"
	"net/http"
)

func main() {
	show := flag.Bool("show", false, "a bool")
	stop := flag.Bool("stop", false, "a bool")
	port := flag.Int("port", 12346, "program port")
	flag.Parse()

	if *show {
		resp, err := http.Get(fmt.Sprintf("http://localhost:%d/show", *port))
		if err != nil {
			// handle error
		}
		resp.Body.Close()
	}

	if *stop {
		resp, err := http.Get(fmt.Sprintf("http://localhost:%d/quit", *port))
		if err != nil {
			// handle error
		}
		resp.Body.Close()
	}
}
