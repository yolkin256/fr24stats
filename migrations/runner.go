package migrations

import (
	"database/sql"
	"os"

	"github.com/pressly/goose/v3"
)

func Run(db *sql.DB, path string) error {
	fs := os.DirFS(path)
	goose.SetBaseFS(fs)

	if err := goose.SetDialect("sqlite3"); err != nil {
		return err
	}

	if err := goose.Up(db, "sql"); err != nil {
		return err
	}
	return nil
}
