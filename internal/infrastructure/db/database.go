package db

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

func ConnectDB() (*sql.DB, error) {
	dbPath, err := getDatabasePath("clipboard.db")
	if err != nil {
		return nil, err
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

	return db, nil
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
