package main

import (
	"flag"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/driver/desktop"

)

func main() {
	no_tray := flag.Bool("i", false, "whether run in background")
	hide := flag.Bool("hide", false, "hide window on launch")
	flag.Parse()
	tray := !*no_tray
	a := app.New()
	w := a.NewWindow("Disk Usage")

	w.Resize(fyne.NewSize(585, 0))
	cons_window(w)

	if tray {
		if desk, ok := a.(desktop.App); ok {
			m := fyne.NewMenu("MyApp",
				fyne.NewMenuItem("Show", func() {
					cons_window(w)
					w.Show()
				}))
			desk.SetSystemTrayMenu(m)
		}
		w.SetCloseIntercept(func() {
			w.Hide()
		})
	}

	create_bindings(w, a, tray)
	create_server(w, a)

	if !*hide {
		w.Show()
	}
	w.CenterOnScreen()
	a.Run()
}
