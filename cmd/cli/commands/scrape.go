package commands

import (
	"errors"
	"fmt"
	"fr24stats/internal/entity"
	"fr24stats/internal/usecase/scrape"
	"github.com/genericplatform/flightradar24sdk"
	"github.com/spf13/cobra"
	"log"
)

func NewScrapeCmd(cfg AppConfig) *cobra.Command {
	var airline string
	cmd := &cobra.Command{
		Use:   "scrape",
		Short: "Получение данных о полётах",
		RunE: func(cmd *cobra.Command, _ []string) error {
			if airline == "" {
				return errors.New("не указан код авиалинии")
			}

			db, err := initDB(cfg)
			if err != nil {
				return fmt.Errorf("ошибка инициализации БД: %w", err)
			}
			defer db.Close()

			log.Println("[INFO] Запуск сохранения данных о полётах в БД")

			handler := scrape.NewHandler(
				flightradar24sdk.NewAPI(
					flightradar24sdk.WithDebug(Verbose),
					flightradar24sdk.WithLogger(createAPIClientLogger()),
				),
				entity.NewFlightRepository(db),
			)

			return handler.Handle(cmd.Context(), scrape.Cmd{Airline: airline})
		},
	}

	cmd.Flags().StringVarP(&airline, "airline", "", "", "Код авиалинии (например, AFL)")

	return cmd
}
