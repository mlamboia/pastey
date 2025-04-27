package repository

import (
	"database/sql"
	"pastey/internal/entity"
)

type ClipboardRepository interface {
	Save(item entity.ClipboardItem) error
	Delete(id int) error
	TogglePin(id int) error
	GetHistory(limit int, offset int) ([]entity.ClipboardItem, error)
}

type clipboardRepository struct {
	db *sql.DB
}

func NewClipboardRepository(db *sql.DB) ClipboardRepository {
	return &clipboardRepository{db}
}

func (r *clipboardRepository) Save(item entity.ClipboardItem) error {
	_, err := r.db.Exec(`
		INSERT INTO clipboard_items (content, timestamp) 
		VALUES (?, ?)
	`, item.Content, item.Timestamp)

	return err
}

func (r *clipboardRepository) TogglePin(id int) error {
	_, err := r.db.Exec(`
		UPDATE clipboard_items
		SET pinned = NOT pinned
		WHERE id = ?
	`, id)

	return err
}

func (r *clipboardRepository) Delete(id int) error {
	_, err := r.db.Exec(`
		DELETE FROM clipboard_items
		WHERE id = ?
	`, id)

	return err
}

func (r *clipboardRepository) GetHistory(limit int, offset int) ([]entity.ClipboardItem, error) {
	rows, err := r.db.Query(`
		SELECT id, pinned, content, timestamp 
		FROM clipboard_items 
		ORDER BY timestamp DESC 
		LIMIT ? OFFSET ?`, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []entity.ClipboardItem
	for rows.Next() {
		var item entity.ClipboardItem
		if err := rows.Scan(&item.ID, &item.Pinned, &item.Content, &item.Timestamp); err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, nil
}
