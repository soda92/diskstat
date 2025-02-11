package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"

	"github.com/ricochet2200/go-disk-usage/du"

	"golang.org/x/sys/windows"
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

func get_disk_name(v string) string {
	volumePath := v // Replace with the desired volume path
	var volumeNameBuffer [windows.MAX_PATH + 1]uint16
	var fileSystemNameBuffer [windows.MAX_PATH + 1]uint16
	var volumeSerialNumber uint32
	var maxComponentLength uint32
	var fileSystemFlags uint32

	err := windows.GetVolumeInformation(
		windows.StringToUTF16Ptr(volumePath),
		&volumeNameBuffer[0],
		uint32(len(volumeNameBuffer)),
		&volumeSerialNumber,
		&maxComponentLength,
		&fileSystemFlags,
		&fileSystemNameBuffer[0],
		uint32(len(fileSystemNameBuffer)),
	)

	if err != nil {
		log.Fatalf("Error getting volume information: %v", err)
	}

	volumeName := windows.UTF16ToString(volumeNameBuffer[:])
	return fmt.Sprintf("%s (%c:)", volumeName, v[0])
}

func cons_window(w fyne.Window) {
	x := container.New(layout.NewVBoxLayout())

	disks := get_disks()
	for _, v := range disks {
		x.Add(widget.NewLabel(get_disk_name(v)))
		x.Add(widget.NewLabel(get_usage(v)))
	}

	w.SetContent(x)
}

func create_bindings(w fyne.Window, a fyne.App, tray bool) {
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
			if tray {
				w.Hide()
			} else {
				a.Quit()
			}
		}
	})
}

func create_server(w fyne.Window, a fyne.App) {
	server := &http.Server{
		Addr: "localhost:12347",
	}

	http.HandleFunc("/show", func(rw http.ResponseWriter, r *http.Request) {
		w.Show()
	})

	http.HandleFunc("/quit", func(rw http.ResponseWriter, r *http.Request) {
		a.Quit()
	})

	go server.ListenAndServe()
}

func main() {
	no_tray := flag.Bool("i", false, "whether run in background")
	flag.Parse()
	tray := !*no_tray
	a := app.New()
	w := a.NewWindow("Disk Usage")
	w.CenterOnScreen()

	w.Resize(fyne.NewSize(585, 444))
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

	w.ShowAndRun()
}
