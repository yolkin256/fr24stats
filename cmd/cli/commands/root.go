package commands

import (
	"github.com/spf13/cobra"
)

type Config struct {
	DBFILE string `env:"FR24_STATS_DB" description:"строка подключения к БД"`
}

func NewRootCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "app",
		Short: "Flightradar24 stats service",
	}
}
