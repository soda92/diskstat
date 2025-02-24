package main

import (
	"fmt"
	"net/http"

	"fyne.io/fyne/v2"
)

func CreateServer(mw MyWindow, a fyne.App, port int) {
	server := &http.Server{
		Addr: fmt.Sprintf("localhost:%d", port),
	}

	http.HandleFunc("/show", func(rw http.ResponseWriter, r *http.Request) {
		mw.w.Show()
		mw.RefreshWindow()
	})

	http.HandleFunc("/quit", func(rw http.ResponseWriter, r *http.Request) {
		a.Quit()
	})

	go server.ListenAndServe()
}
