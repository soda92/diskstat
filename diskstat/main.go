package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"

	"flag"
)

func main() {
	hide := flag.Bool("hide", false, "hide window")
	cmd := flag.Bool("cmd", false, "run on foreground - close on window close")
	port := flag.Int("port", 12347, "server port")
	flag.Parse()

	a := app.New()
	w := a.NewWindow("Disk Usage")
	w.Resize(fyne.NewSize(585, 0))

	my_window := InitWindow(w)

	CreateShortcuts(my_window, a, !*cmd)
	CreateServer(my_window, a, *port)

	if !*cmd {
		w.SetCloseIntercept(func() { w.Hide() })
	}
	if !*hide {
		w.Show()
	}
	w.CenterOnScreen()
	a.Run()
}
