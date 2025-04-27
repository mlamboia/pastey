package usecase

import (
	"pastey/internal/entity"
	"pastey/internal/repository"
	"time"
)

type ClipboardUseCase interface {
	GetHistory(limit int, offset int) ([]entity.ClipboardItem, error)
	SaveClipboardContent(content string) error
	TogglePinItem(id int) error
}

type clipboardUseCase struct {
	Repo repository.ClipboardRepository
}

func NewClipboardUseCase(repo repository.ClipboardRepository) ClipboardUseCase {
	return &clipboardUseCase{Repo: repo}
}

func (uc *clipboardUseCase) GetHistory(limit int, offset int) ([]entity.ClipboardItem, error) {
	return uc.Repo.GetHistory(limit, offset)
}

func (uc *clipboardUseCase) SaveClipboardContent(content string) error {
	item := entity.ClipboardItem{
		Content:   content,
		Timestamp: time.Now(),
	}
	return uc.Repo.Save(item)
}

func (uc *clipboardUseCase) TogglePinItem(id int) error {
	return uc.Repo.TogglePin(id)
}
