package commands

import (
	"database/sql"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/sqliteshim"
	"github.com/uptrace/bun/extra/bundebug"
	"log"
)

var Verbose bool

type AppConfig struct {
	DBFile string `env:"FR24_STATS_DB" description:"строка подключения к БД"`
}

func NewRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "app",
		Short: "Flightradar24 stats service",
	}
	cmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "подробный вывод")
	return cmd
}

func initDB(cfg AppConfig) (*bun.DB, error) {
	sqldb, err := sql.Open(sqliteshim.ShimName, fmt.Sprintf("file:%s", cfg.DBFile))
	if err != nil {
		return nil, fmt.Errorf("ошибка открытия БД: %w", err)
	}

	db := bun.NewDB(sqldb, sqlitedialect.New())
	if Verbose {
		db.AddQueryHook(bundebug.NewQueryHook(
			bundebug.WithVerbose(true),
		))
	}
	return db, nil
}

type apiLogger struct{}

func (l *apiLogger) Errorf(format string, v ...interface{}) {
	log.Printf("[ERROR]"+format, v)
}
func (l *apiLogger) Warnf(format string, v ...interface{}) {
	log.Printf("[WARN]"+format, v)
}
func (l *apiLogger) Debugf(format string, v ...interface{}) {
	log.Printf("[DEBUG]"+format, v)
}

func createAPIClientLogger() *apiLogger {
	return &apiLogger{}
}
