package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	a := app.New()
	w := a.NewWindow("Disk Usage")

	w.Resize(fyne.NewSize(585, 0))
	cons_window(w)

	create_bindings(w, a, true)
	create_server(w, a)

	w.SetCloseIntercept(func() {
		w.Hide()
	})
	w.CenterOnScreen()
	a.Run()
}
