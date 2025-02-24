package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func InitWindow(w fyne.Window) {
	x := container.New(layout.NewVBoxLayout())

	usages := AllDiskUsage()

	for _, v := range usages {
		name := v.disk_name
		x.Add(widget.NewLabel(name))
		progress := v.PBar()
		x.Add(progress)
		label := widget.NewLabel(v.Label())
		x.Add(label)
	}

	w.SetContent(x)
}
