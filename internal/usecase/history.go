package usecase

import (
	"pastey/internal/entity"
	"pastey/internal/repository"
)

type HistoryUsecase interface {
	GetHistory(limit, offset int) ([]entity.ClipboardItem, error)
	DeleteItem(id int) error
	PinItem(id int) error
}

type historyUsecase struct {
	repo repository.ClipboardRepository
}

func NewHistoryUsecase(repo repository.ClipboardRepository) HistoryUsecase {
	return &historyUsecase{repo: repo}
}

func (uc *historyUsecase) GetHistory(limit, offset int) ([]entity.ClipboardItem, error) {
	items, err := uc.repo.GetHistory(limit, offset)
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (uc *historyUsecase) DeleteItem(id int) error {
	err := uc.repo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func (uc *historyUsecase) PinItem(id int) error {
	err := uc.repo.TogglePin(id)
	if err != nil {
		return err
	}
	return nil
}
