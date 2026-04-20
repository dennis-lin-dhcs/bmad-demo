package store

import (
    "database/sql"
    "os"
    "path/filepath"

    _ "modernc.org/sqlite"
)

func Open() (*sql.DB, error) {
    exePath, err := os.Executable()
    if err != nil {
        return nil, err
    }

    dbPath := filepath.Join(filepath.Dir(exePath), "app-template.db")
    db, err := sql.Open("sqlite", dbPath)
    if err != nil {
        return nil, err
    }

    if _, err := db.Exec(`
        CREATE TABLE IF NOT EXISTS llm_settings (
            provider TEXT PRIMARY KEY,
            endpoint TEXT NOT NULL,
            model TEXT NOT NULL,
            api_key TEXT NOT NULL DEFAULT '',
            temperature REAL NOT NULL DEFAULT 0.7,
            system_prompt TEXT NOT NULL DEFAULT '',
            updated_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP
        )
    `); err != nil {
        _ = db.Close()
        return nil, err
    }

    return db, nil
}
