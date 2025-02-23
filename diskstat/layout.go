package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/ricochet2200/go-disk-usage/du"
)

func cons_window(w fyne.Window) {
	x := container.New(layout.NewVBoxLayout())

	disks := get_disks()
	for _, v := range disks {
		name := get_disk_name(v)
		if name == "Google Drive (G:)" {
			continue
		}
		x.Add(widget.NewLabel(name))
		progress := widget.NewProgressBar()
		usage := du.NewDiskUsage(v)
		progress.Value = float64(usage.Used())
		progress.Max = float64(usage.Size())
		if float64(usage.Used())/float64(usage.Size()) > 0.9 {
			// progress.Theme()
			// TODO add red color
		}
		x.Add(progress)
		x.Add(widget.NewLabel(get_usage(v)))
	}

	w.SetContent(x)
}
