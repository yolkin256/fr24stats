package main

import (
	"fr24stats/cmd/cli/commands"
	"github.com/fatih/color"
	"github.com/go-pkgz/lgr"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

func main() {
	setupLog()
	if err := godotenv.Load(); err != nil {
		log.Fatalf("[FATAL] Ошибка загрузки .env файла: %v", err)
	}

	cfg := commands.Config{
		DBFILE: os.Getenv("FR24_STATS_DB"),
	}

	rootCmd := commands.NewRootCmd()
	rootCmd.AddCommand(commands.NewMigrationCmd())
	rootCmd.AddCommand(commands.NewScrapeCmd(cfg))

	if err := rootCmd.Execute(); err != nil {
		log.Printf("[ERROR] Ошибка выполнения команды: %v", err)
	} else {
		log.Printf("[INFO] Операция успешно завершена\n")
	}
}

func setupLog() {
	logOpts := []lgr.Option{lgr.Debug, lgr.LevelBraces}

	colorizer := lgr.Mapper{
		ErrorFunc:  func(s string) string { return color.New(color.FgHiRed).Sprint(s) },
		WarnFunc:   func(s string) string { return color.New(color.FgRed).Sprint(s) },
		InfoFunc:   func(s string) string { return color.New(color.FgYellow).Sprint(s) },
		DebugFunc:  func(s string) string { return color.New(color.FgWhite).Sprint(s) },
		CallerFunc: func(s string) string { return color.New(color.FgBlue).Sprint(s) },
		TimeFunc:   func(s string) string { return color.New(color.FgCyan).Sprint(s) },
	}
	logOpts = append(logOpts, lgr.Map(colorizer))

	lgr.SetupStdLogger(logOpts...)
	lgr.Setup(logOpts...)
}
