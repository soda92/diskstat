package main

import (
	"fmt"

	"fyne.io/fyne/v2"
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

type MyWindow struct {
	w      fyne.Window
	usages []disk_usage
}

func AllDiskUsage() []disk_usage {
	disks := get_disks()
	var ret []disk_usage
	for _, v := range disks {
		var d disk_usage
		d.disk_path = v
		d.disk_name = get_disk_name(v)
		if d.disk_name == "Google Drive (G:)" {
			continue
		}

		usage := du.NewDiskUsage(v)
		d.used = float64(usage.Used()) / GB
		d.total = float64(usage.Size()) / GB

		d.is_new = false
		d.is_removed = false
		d.old_used = d.used

		ret = append(ret, d)
	}
	return ret
}

func RefreshDiskUsage(old []disk_usage) []disk_usage {
	new := AllDiskUsage()
	disks := DiskPaths(new)

	for _, d := range disks {
		if DiskNotIn(old, d) {
			index := DiskIndex(new, d)
			new[index].is_new = true
		}
	}

	for _, d := range DiskPaths(old) {
		if DiskNotIn(new, d) {
			index := DiskIndex(old, d)
			val := old[index]
			val.is_removed = true

			index = FindIndex(DiskPaths(new), val.disk_path)
			new = InsertOrdered(new, val, index)
		} else if DiskNotIn(old, d) {
			index := DiskIndex(new, d)
			new[index].is_new = true
		} else {
			index1 := DiskIndex(old, d)
			index2 := DiskIndex(new, d)
			new[index2].old_used = old[index1].used
		}
	}

	return new
}

func (d *disk_usage) Label() string {
	str := fmt.Sprintf("%.1fGB free of %.0fGB",
		d.total-d.used, d.total)
	if d.old_used != d.used {
		diff := d.used - d.old_used
		diffMB := diff * 1024
		percent := diff / d.total * 100
		sign := ""
		if diff > 0 {
			sign = "+"
		}
		if diffMB > 1 || diffMB < -1 {
			if percent < 1 && percent > -1 {
				str2 := fmt.Sprintf("(%s%.0fMB)", sign, diffMB)
				str += str2
			} else {
				str2 := fmt.Sprintf("(%s%.1f%s/%.0fMB)", sign, percent, "%", diffMB)
				str += str2
			}
		}
	}
	if d.is_new {
		str += "(newly added)"
	}
	if d.is_removed {
		str += "(removed)"
	}
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
