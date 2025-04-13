package usecase

import (
    "pastey/internal/entity"
    "pastey/internal/repository"
    "time"
)

type ClipboardUseCase struct {
    Repo repository.ClipboardRepository
}

func (uc *ClipboardUseCase) GetHistory(limit int, offset int)  ([]entity.ClipboardItem, error) {
    return uc.Repo.GetHistory(limit, offset)
}

func (uc *ClipboardUseCase) SaveClipboardContent(content string) error {
    item := entity.ClipboardItem{
        Content:   content,
        Timestamp: time.Now(),
    }
    return uc.Repo.Save(item)
}

func (uc *ClipboardUseCase) TogglePinItem(id int) error {
    return uc.Repo.TogglePin(id)
}
