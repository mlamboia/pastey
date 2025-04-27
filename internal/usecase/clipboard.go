package usecase

import (
	"pastey/internal/entity"
	"pastey/internal/repository"
	"time"
)

type ClipboardUsecase interface {
	GetByContent(content string) (entity.ClipboardItem, error)
	UpdateTimestamp(id int) error
	SaveClipboardContent(content string) error
	TogglePinItem(id int) error
}

type clipboardUsecase struct {
	Repo repository.ClipboardRepository
}

func NewClipboardUsecase(repo repository.ClipboardRepository) ClipboardUsecase {
	return &clipboardUsecase{Repo: repo}
}

func (uc *clipboardUsecase) GetByContent(content string) (entity.ClipboardItem, error) {
	return uc.Repo.GetByContent(content)
}

func (uc *clipboardUsecase) UpdateTimestamp(id int) error {
	return uc.Repo.UpdateTimestamp(id)
}

func (uc *clipboardUsecase) SaveClipboardContent(content string) error {
	item := entity.ClipboardItem{
		Content:   content,
		Timestamp: time.Now(),
	}
	return uc.Repo.Save(item)
}

func (uc *clipboardUsecase) TogglePinItem(id int) error {
	return uc.Repo.TogglePin(id)
}
