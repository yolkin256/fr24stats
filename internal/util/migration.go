package util

import (
	"database/sql"
	"github.com/pressly/goose/v3"
	"os"
	"path/filepath"
)

func RunMigrations(db *sql.DB, path string) error {
	p1, p2 := filepath.Split(path)
	fs := os.DirFS(p1)
	goose.SetBaseFS(fs)

	if err := goose.SetDialect("sqlite3"); err != nil {
		return err
	}

	if err := goose.Up(db, p2); err != nil {
		return err
	}
	return nil
}
