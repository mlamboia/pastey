package main

import (
    "pastey/internal/infrastructure/clipboard"
    "pastey/internal/infrastructure/db"
    "pastey/internal/usecase"
    "fmt"
    "time"
)

func main() {
    repo, err := db.NewSqliteRepository("clipboard.db")
    if err != nil {
        panic(err)
    }

    useCase := usecase.ClipboardUseCase{Repo: repo}
    var last string

    for {
        text, err := clipboard.ReadClipboard()
		fmt.Println(text)
        if err == nil && text != "" && text != last {
            err = useCase.SaveClipboardContent(text)

            if err != nil {
                fmt.Println("Save failed:", err)
            } else {
                fmt.Println("Copied:", text)
                last = text
            }
        }

        time.Sleep(1 * time.Second)
    }
}
