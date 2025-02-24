package main

import (
	"fmt"

	"fyne.io/fyne/v2/widget"

	"github.com/ricochet2200/go-disk-usage/du"
)

type disk_usage struct {
	disk_name  string
	disk_path  string
	used       float64
	total      float64
	old_used   float64
	is_new     bool
	is_removed bool
}

func AllDiskUsage() []disk_usage {
	disks := get_disks()
	var ret []disk_usage
	for _, v := range disks {
		var d disk_usage
		d.disk_path = v
		d.disk_name = get_disk_name(v)
		if d.disk_name == "Google Drive" {
			continue
		}

		usage := du.NewDiskUsage(v)
		d.used = float64(usage.Used()) / GB
		d.total = float64(usage.Size()) / GB
		ret = append(ret, d)
	}
	return ret
}

func (d *disk_usage) Label() string {
	str := fmt.Sprintf("%.1fGB free of %.0fGB",
		d.total-d.used, d.total)
	return str
}

func (d *disk_usage) PBar() *widget.ProgressBar {
	progress := widget.NewProgressBar()
	usage := du.NewDiskUsage(d.disk_path)
	progress.Value = float64(usage.Used())
	progress.Max = float64(usage.Size())
	if float64(usage.Used())/float64(usage.Size()) > 0.9 {
		// progress.Theme()
		// TODO add red color
	}
	return progress
}
