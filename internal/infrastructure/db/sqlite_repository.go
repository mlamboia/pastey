package db

import (
    "pastey/internal/entity"
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
)

type SqliteRepository struct {
    DB *sql.DB
}

func NewSqliteRepository(dbPath string) (*SqliteRepository, error) {
    db, err := sql.Open("sqlite3", dbPath)
    if err != nil {
        return nil, err
    }

    createTable := `
		CREATE TABLE IF NOT EXISTS clipboard_items (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			pinned BOOLEAN DEFAULT FALSE,
			content TEXT,
			timestamp DATETIME
		);
    `
    _, err = db.Exec(createTable)
    if err != nil {
        return nil, err
    }

    return &SqliteRepository{DB: db}, nil
}

func (r *SqliteRepository) Save(item entity.ClipboardItem) error {
    _, err := r.DB.Exec(`
		INSERT INTO clipboard_items (content, timestamp) 
		VALUES (?, ?)
	`, item.Content, item.Timestamp)

    return err
}

func (r *SqliteRepository) TogglePin(id int) error {
    _, err := r.DB.Exec(`
		UPDATE clipboard_items
		SET pinned = NOT pinned
		WHERE Id = ?
	`, id)

    return err
}

func (r *SqliteRepository) GetHistory(limit int, offset int) ([]entity.ClipboardItem, error) {
    rows, err := r.DB.Query(`
		SELECT id, pinned, content, timestamp 
		FROM clipboard_items 
		ORDER BY timestamp DESC 
		LIMIT ?
		OFFSET ?
	`, limit, offset)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var items []entity.ClipboardItem
    for rows.Next() {
        var i entity.ClipboardItem
        if err := rows.Scan(&i.ID, &i.Content, &i.Timestamp); err != nil {
            return nil, err
        }
        items = append(items, i)
    }

    return items, nil
}
