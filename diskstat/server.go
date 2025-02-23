package main

import (
	"net/http"

	"fyne.io/fyne/v2"
)

func create_server(w fyne.Window, a fyne.App) {
	server := &http.Server{
		Addr: "localhost:12347",
	}

	http.HandleFunc("/show", func(rw http.ResponseWriter, r *http.Request) {
		w.Show()
	})

	http.HandleFunc("/quit", func(rw http.ResponseWriter, r *http.Request) {
		a.Quit()
	})

	go server.ListenAndServe()
}