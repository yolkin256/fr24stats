package commands

import (
	"database/sql"
	"errors"
	"fmt"
	"fr24stats/migrations"
	"log"

	"github.com/spf13/cobra"
)

func NewMigrationCmd() *cobra.Command {
	var file string
	cmd := &cobra.Command{
		Use:   "migration",
		Short: "Раскатка миграций БД",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("[INFO] Запуск миграций")

			if file == "" {
				return errors.New("не указан путь к БД")
			}
			db, err := sql.Open("sqlite3", file)
			if err != nil {
				return fmt.Errorf("ошибка открытия БД: %w", err)
			}

			return migrations.Run(db)
		},
	}

	cmd.Flags().StringVarP(&file, "file", "", "", "Путь к файлу базе данных SQLite (файл создастся, если ещё не существует)")

	return cmd
}
