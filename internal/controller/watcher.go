package controller

import (
	"log"
	"pastey/internal/infrastructure/clipboard"
	"pastey/internal/usecase"
	"strings"
	"time"
)

func WatchClipboard(uc *usecase.ClipboardUseCase, interval time.Duration) {
	var last string

	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for range ticker.C {
		text, err := clipboard.ReadClipboard()
		if err != nil {
			log.Println("Clipboard read error:", err)
			continue
		}

		text = strings.TrimSpace(text)
		if text != "" && text != last {
			err := uc.SaveClipboardContent(text)
			if err != nil {
				log.Println("Save error:", err)
			} else {
				last = text
				log.Println("Copied:", text)
			}
		}
	}
}
