package main

import (
	"fmt"
	"net/http"

	"fyne.io/fyne/v2"
)

func CreateServer(w fyne.Window, a fyne.App, port int) {
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
