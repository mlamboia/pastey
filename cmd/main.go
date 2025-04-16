package main

import (
	"pastey/internal/controller"
	"pastey/internal/infrastructure/db"
	"pastey/internal/usecase"
	"time"
)

func main() {
	repo, err := db.NewSqliteRepository("clipboard.db")
	if err != nil {
		panic(err)
	}

	clipboardUseCase := &usecase.ClipboardUseCase{Repo: repo}
	controller.WatchClipboard(clipboardUseCase, 1*time.Second)
}
