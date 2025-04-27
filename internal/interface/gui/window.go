package gui

import (
	"pastey/internal/interface/controller"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func SetupMainWindow(app fyne.App) fyne.Window {
	w := app.NewWindow("Pastey")
	ctrl := controller.NewShortcutController(w)
	go ctrl.GlobalShortcutsListener()

	w.SetContent(container.NewVBox(
		widget.NewLabel("Press Ctrl + Alt + H to minimize"),
	))
	w.Resize(fyne.NewSize(400, 200))
	w.CenterOnScreen()

	return w
}
