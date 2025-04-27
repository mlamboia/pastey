package gui

import (
	"pastey/internal/interface/gui/screen"
	"pastey/internal/usecase"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

type Screen interface {
	Build() fyne.CanvasObject
}

type AppWindow struct {
	app            fyne.App
	window         fyne.Window
	currentScreen  Screen
	historyUsecase usecase.HistoryUsecase
}

func NewAppWindow(historyUsecase usecase.HistoryUsecase) *AppWindow {
	app := app.New()
	window := app.NewWindow("Pastey")

	initialScreen := screen.NewHistoryScreen(historyUsecase)

	return &AppWindow{
		app:            app,
		window:         window,
		currentScreen:  initialScreen,
		historyUsecase: historyUsecase,
	}
}

func (a *AppWindow) Build() {
	a.ShowHistory()
	a.window.Resize(fyne.NewSize(600, 400))
}

func (a *AppWindow) ShowHistory() {
	a.currentScreen = screen.NewHistoryScreen(a.historyUsecase)
	a.window.SetContent(a.currentScreen.Build())
}

func (a *AppWindow) Reload() {
	fyne.DoAndWait(func() {
		a.window.SetContent(a.currentScreen.Build())
	})
}

// func (a *AppWindow) ShowClipboard() {
// 	clipboardScreen := screen.NewClipboardWindow(a.historyUsecase)
// 	a.window.SetContent(clipboardScreen.Build())
// }

func (a *AppWindow) Run() {
	a.window.ShowAndRun()
}
