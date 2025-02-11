package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Disk Usage")

	w.Resize(fyne.NewSize(585, 444))

	objects := []widget.Label{}
	objects = append(objects, *widget.NewLabel("aa"))
	x := container.New(layout.NewVBoxLayout())
	x.Add(&objects[0])

	w.SetContent(x)
	w.ShowAndRun()
}
