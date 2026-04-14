package main

import (
	"os"

	"github.com/charmbracelet/crush/internal/app"
	"github.com/charmbracelet/crush/internal/config"
	"github.com/charmbracelet/log"
)

func main() {
	// Load configuration from environment and config files
	cfg, err := config.Load()
	if err != nil {
		log.Error("failed to load configuration", "error", err)
		os.Exit(1)
	}

	// Initialize and run the application
	application, err := app.New(cfg)
	if err != nil {
		log.Error("failed to initialize application", "error", err)
		os.Exit(1)
	}

	if err := application.Run(); err != nil {
		log.Error("application exited with error", "error", err)
		os.Exit(1)
	}
}
