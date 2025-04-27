package screen

import (
	"pastey/internal/usecase"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type HistoryScreen struct {
	historyUsecase usecase.HistoryUsecase
	items          []ClipboardItemView
	list           *widget.List
}

type ClipboardItemView struct {
	Content string
}

func NewHistoryScreen(historyUsecase usecase.HistoryUsecase) *HistoryScreen {
	screen := &HistoryScreen{
		historyUsecase: historyUsecase,
	}
	screen.buildList()
	return screen
}

func (h *HistoryScreen) buildList() {
	h.list = widget.NewList(
		func() int {
			return len(h.items)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("Template")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(h.items[i].Content)
		},
	)
}

func (h *HistoryScreen) Build() fyne.CanvasObject {
	h.updateList()
	return container.NewBorder(nil, nil, nil, nil, h.list)
}

func (h *HistoryScreen) updateList() {
	data, err := h.historyUsecase.GetHistory(50, 0)
	if err != nil {
		return
	}
	h.items = make([]ClipboardItemView, len(data))
	for i, item := range data {
		h.items[i] = ClipboardItemView{
			Content: item.Content,
		}
	}
	if h.list != nil {
		h.list.Refresh()
	}
}
