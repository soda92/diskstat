package main

import (
	"fmt"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"

	"github.com/ricochet2200/go-disk-usage/du"
)

var KB = float64(1024)
var GB = KB * KB * KB

func pathExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func get_disks() []string {
	x := []string{}
	for char := 'A'; char <= 'Z'; char++ {
		path := fmt.Sprintf("%c%s", char, ":\\")
		if pathExists(path) {
			x = append(x, path)
		}
	}
	return x
}

func get_usage(d string) string {
	usage := du.NewDiskUsage(d)

	str := fmt.Sprintf("%.1fGB free of %.0fGB",
		float64(usage.Free())/GB, float64(usage.Size())/GB)
	return str
}

func cons_window(w fyne.Window) {
	x := container.New(layout.NewVBoxLayout())

	disks := get_disks()
	for _, v := range disks {
		x.Add(widget.NewLabel(v))
		x.Add(widget.NewLabel(get_usage(v)))
	}

	w.SetContent(x)
}

func main() {
	a := app.New()
	w := a.NewWindow("Disk Usage")
	w.CenterOnScreen()

	w.Resize(fyne.NewSize(585, 444))
	cons_window(w)

	if desk, ok := a.(desktop.App); ok {
		m := fyne.NewMenu("MyApp",
			fyne.NewMenuItem("Show", func() {
				// log.Println("Tapped show")
				w.Show()
			}))
		desk.SetSystemTrayMenu(m)
	}

	CtrlR := &desktop.CustomShortcut{KeyName: fyne.KeyR, Modifier: fyne.KeyModifierControl}
	w.Canvas().AddShortcut(CtrlR, func(shortcut fyne.Shortcut) {
		cons_window(w)
	})

	CtrlQ := &desktop.CustomShortcut{KeyName: fyne.KeyQ, Modifier: fyne.KeyModifierControl}
	w.Canvas().AddShortcut(CtrlQ, func(shortcut fyne.Shortcut) {
		a.Quit()
	})

	w.Canvas().SetOnTypedKey(func(key *fyne.KeyEvent) {
		if key.Name == fyne.KeyEscape {
			w.Hide()
		}
	})

	w.SetCloseIntercept(func() {
		w.Hide()
	})

	w.ShowAndRun()
}
