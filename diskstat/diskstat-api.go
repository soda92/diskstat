package main

import "net/http"
import "flag"

func main() {
	show := flag.Bool("show", false, "a bool")
	stop := flag.Bool("stop", false, "a bool")
	flag.Parse()
	if *show {
		resp, err := http.Get("http://localhost:12346/show")
		if err != nil {
			// handle error
		}
		resp.Body.Close()
	}

	if *stop {
		resp, err := http.Get("http://localhost:12346/quit")
		if err != nil {
			// handle error
		}
		resp.Body.Close()
	}
}
