package main

import (
	"pastey/internal/controller"
	"pastey/internal/infrastructure/db"
	"pastey/internal/interface/gui"
	"pastey/internal/repository"
	"pastey/internal/usecase"
	"time"
)

func main() {
	database, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	clipboardRepo := repository.NewClipboardRepository(database)
	clipboardUsecase := usecase.NewClipboardUsecase(clipboardRepo)
	historyUsecase := usecase.NewHistoryUsecase(clipboardRepo)

	w := gui.NewAppWindow(historyUsecase)

	go controller.WatchClipboard(clipboardUsecase, 1*time.Second, w.Reload)

	w.Build()
	w.Run()
}
