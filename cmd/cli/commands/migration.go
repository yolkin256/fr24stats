package commands

import (
	"database/sql"
	"fmt"
	"fr24stats/internal/util"
	"log"

	"github.com/spf13/cobra"
)

func NewMigrationCmd(cfg AppConfig) *cobra.Command {
	var path string
	cmd := &cobra.Command{
		Use:   "migrate",
		Short: "Раскатка миграций БД",
		RunE: func(_ *cobra.Command, _ []string) error {
			log.Println("[INFO] Запуск миграций")

			db, err := sql.Open("sqlite3", cfg.DBFile)
			if err != nil {
				return fmt.Errorf("ошибка открытия БД: %w", err)
			}

			return util.RunMigrations(db, path)
		},
	}

	cmd.Flags().StringVarP(&path, "path", "", "", "Путь к директории с миграциями")

	return cmd
}
