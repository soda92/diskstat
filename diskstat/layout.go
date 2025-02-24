package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func InitWindow(w fyne.Window) MyWindow {
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
	var my_window MyWindow
	my_window.w = w
	my_window.usages = usages
	return my_window
}

func (m *MyWindow) RefreshWindow() {
	x := container.New(layout.NewVBoxLayout())

	usages := RefreshDiskUsage(m.usages)

	for _, v := range usages {
		name := v.disk_name
		x.Add(widget.NewLabel(name))
		progress := v.PBar()
		x.Add(progress)
		label := widget.NewLabel(v.Label())
		x.Add(label)
	}
	m.usages = CleanupUsage(usages)
	m.w.SetContent(x)
}
