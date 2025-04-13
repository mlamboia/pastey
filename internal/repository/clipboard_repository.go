package repository

import "pastey/internal/entity"

type ClipboardRepository interface {
    Save(item entity.ClipboardItem) error
    TogglePin(id int) error
    GetHistory(limit int, offset int) ([]entity.ClipboardItem, error)
}
