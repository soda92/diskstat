package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/desktop"
)

func CreateShortcuts(mw MyWindow, a fyne.App, tray bool) {
	w := mw.w
	CtrlR := &desktop.CustomShortcut{KeyName: fyne.KeyR, Modifier: fyne.KeyModifierControl}
	w.Canvas().AddShortcut(CtrlR, func(shortcut fyne.Shortcut) {
		mw.RefreshWindow()
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
