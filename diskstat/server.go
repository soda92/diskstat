package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"net/http"
)

func create_server(w fyne.Window, a fyne.App, port int) {
	server := &http.Server{
		Addr: fmt.Sprintf("localhost:%d", port),
	}

	http.HandleFunc("/show", func(rw http.ResponseWriter, r *http.Request) {
		w.Show()
	})

	http.HandleFunc("/quit", func(rw http.ResponseWriter, r *http.Request) {
		a.Quit()
	})

	go server.ListenAndServe()
}
