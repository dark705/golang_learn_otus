package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/dark705/otus/hw08/internal/calendar/calendar"
	"github.com/dark705/otus/hw08/internal/config"
	"github.com/dark705/otus/hw08/internal/logger"
	"github.com/dark705/otus/hw08/internal/storage"
)

func main() {
	var cFile string
	flag.StringVar(&cFile, "config", "config/config.yaml", "Config file")
	flag.Parse()
	if cFile == "" {
		fmt.Println("Not set config file")
		os.Exit(0)
	}

	config, err := config.ReadFromFile(cFile)
	if err != nil {
		_, _ = fmt.Fprint(os.Stderr, err)
		os.Exit(2)
	}

	logger := logger.GetLogger(config)

	inMemory := storage.InMemory{}
	inMemory.Init()
	_ = calendar.Calendar{Config: config, Storage: &inMemory, Logger: logger}
}
