package main

import (
	"fmt"
	"os"

	"github.com/dark705/otus/hw08/internal/calendar/calendar"
	"github.com/dark705/otus/hw08/internal/config"
	"github.com/dark705/otus/hw08/internal/logger"
	"github.com/dark705/otus/hw08/internal/storage"
)

func main() {

	config, err := config.ReadFromFile("config/config.yaml")
	if err != nil {
		_, _ = fmt.Fprint(os.Stderr, err)
		os.Exit(2)
	}
	logger := logger.GetLogger(config)

	inMemory := storage.InMemory{}
	inMemory.Init()
	_ = calendar.Calendar{Config: config, Storage: &inMemory, Logger: logger}
}
