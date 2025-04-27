package main

import (
	"fmt"
	"pastey/internal/controller"
	"pastey/internal/infrastructure/db"
	"pastey/internal/interface/gui"
	"pastey/internal/repository"
	"pastey/internal/usecase"
	"time"

	"fyne.io/fyne/v2/app"
)

func main() {
	database, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}
	fmt.Println(database)

	clipboardRepo := repository.NewClipboardRepository(database)
	clipboardUseCase := usecase.NewClipboardUseCase(clipboardRepo)
	// historyUseCase := usecase.NewHistoryUseCase(clipboardRepo)
	go controller.WatchClipboard(clipboardUseCase, 1*time.Second)
	a := app.New()
	w := gui.SetupMainWindow(a)
	w.ShowAndRun()
}
