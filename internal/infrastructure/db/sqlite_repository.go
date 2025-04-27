package db

import (
	"database/sql"
	"fmt"
	"os"
	"pastey/internal/entity"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

type SqliteRepository struct {
	DB *sql.DB
}

func NewSqliteRepository(dbName string) (*SqliteRepository, error) {
	dbPath, err := getDatabasePath(dbName)
	if err != nil {
		return nil, fmt.Errorf("unable to determine database path: %v", err)
	}

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	createTableSQL := `
		CREATE TABLE IF NOT EXISTS clipboard_items (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			pinned BOOLEAN DEFAULT FALSE,
			content TEXT,
			timestamp DATETIME
		);
	`

	_, err = db.Exec(createTableSQL)
	if err != nil {
		return nil, err
	}

	return &SqliteRepository{DB: db}, nil
}

func getDatabasePath(dbName string) (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("unable to get user home directory: %v", err)
	}

	var dbDir string
	if isWindows() {
		appData := os.Getenv("APPDATA")
		if appData == "" {
			return "", fmt.Errorf("APPDATA environment variable is not set")
		}
		dbDir = filepath.Join(appData, "Pastey")
	} else {
		dbDir = filepath.Join(homeDir, ".pastey")
	}

	if err := os.MkdirAll(dbDir, 0755); err != nil {
		return "", fmt.Errorf("failed to create database directory: %v", err)
	}

	return filepath.Join(dbDir, dbName), nil
}

func isWindows() bool {
	return filepath.Separator == '\\'
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
		WHERE id = ?
	`, id)

	return err
}

func (r *SqliteRepository) GetHistory(limit int, offset int) ([]entity.ClipboardItem, error) {
	rows, err := r.DB.Query(`
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
