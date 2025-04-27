package controller

import (
	"log"
	"pastey/internal/infrastructure/clipboard"
	"pastey/internal/usecase"
	"strings"
	"time"
)

func WatchClipboard(uc usecase.ClipboardUsecase, interval time.Duration, callback func()) {
	var last string

	firstInteraction := true
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for range ticker.C {
		if firstInteraction {
			firstInteraction = false

			continue
		}

		text, err := clipboard.ReadClipboard()
		if err != nil {
			log.Println("Clipboard read error:", err)
			continue
		}

		text = strings.TrimSpace(text)
		if text != "" && text != last {
			existingItem, err := uc.GetByContent(text)
			if err != nil {
				continue
			}
			if existingItem.ID != 0 {
				continue
			}

			err = uc.SaveClipboardContent(text)
			if err != nil {
				log.Println("Save error:", err)
			} else {
				last = text
				callback()
				log.Println("Copied:", text)
			}
		}
	}
}
