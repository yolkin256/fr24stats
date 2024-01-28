package commands

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"fr24stats/internal/entity"
	"fr24stats/internal/usecase/scrape"
	"github.com/genericplatform/flightradar24sdk"
	"log"

	"github.com/spf13/cobra"
)

func NewScrapeCmd(cfg Config) *cobra.Command {
	var airline string
	cmd := &cobra.Command{
		Use:   "scrape",
		Short: "Получение данных о полётах",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.Background()

			log.Println("[INFO] Запуск сохранения данных о полётах в БД")

			if airline == "" {
				return errors.New("не указан код авиалинии")
			}
			db, err := sql.Open("sqlite3", cfg.DBFILE)
			if err != nil {
				return fmt.Errorf("ошибка открытия БД: %w", err)
			}
			defer db.Close()

			handler := scrape.NewHandler(
				flightradar24sdk.NewAPI(nil),
				entity.NewFlightRepository(db),
			)

			return handler.Handle(ctx, scrape.Cmd{Airline: airline})
		},
	}

	cmd.Flags().StringVarP(&airline, "airline", "", "", "Код авиалинии (например, AFL)")

	return cmd
}
