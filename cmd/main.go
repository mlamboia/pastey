package main

import (
	"pastey/internal/controller"
	"pastey/internal/infrastructure/db"
	"pastey/internal/interface/gui"
	"pastey/internal/usecase"
	"time"

	"fyne.io/fyne/v2/app"
)

func main() {
	repo, err := db.NewSqliteRepository("clipboard.db")
	if err != nil {
		panic(err)
	}

	clipboardUseCase := &usecase.ClipboardUseCase{Repo: repo}
	go controller.WatchClipboard(clipboardUseCase, 1*time.Second)
	a := app.New()
	w := gui.SetupMainWindow(a)
	w.ShowAndRun()
}
