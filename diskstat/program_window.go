package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"os"
)

func pathExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func get_disks() []string {
	x := []string{}
	for char := 'a'; char <= 'z'; char++ {
		path := fmt.Sprintf("%c%s", char, ":\\")
		if pathExists(path) {
			x = append(x, path)
		}
	}
	return x
}

func main() {
	a := app.New()
	w := a.NewWindow("Disk Usage")

	w.Resize(fyne.NewSize(585, 444))
	x := container.New(layout.NewVBoxLayout())

	disks := get_disks()
	for _, v := range disks {
		x.Add(widget.NewLabel(v))
	}

	w.SetContent(x)
	w.ShowAndRun()
}
