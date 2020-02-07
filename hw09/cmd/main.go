package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/dark705/otus/hw09/internal/calendar/calendar"
	"github.com/dark705/otus/hw09/internal/config"
	"github.com/dark705/otus/hw09/internal/logger"
	"github.com/dark705/otus/hw09/internal/storage"
	"github.com/dark705/otus/hw09/internal/web"
)

func main() {
	var cFile string
	flag.StringVar(&cFile, "config", "config/config.yaml", "Config file")
	flag.Parse()
	if cFile == "" {
		_, _ = fmt.Fprint(os.Stderr, "Not set config file")
		os.Exit(2)
	}

	conf, err := config.ReadFromFile(cFile)
	if err != nil {
		_, _ = fmt.Fprint(os.Stderr, err)
		os.Exit(2)
	}

	log := logger.GetLogger(conf)
	defer logger.CloseLogFile()

	osSignals := make(chan os.Signal, 1)
	signal.Notify(osSignals, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)

	inMemory := storage.InMemory{}
	inMemory.Init()
	_ = calendar.Calendar{Config: conf, Storage: &inMemory, Logger: log}

	go web.RunServer(conf, log)

	log.Info("Got %v. Exit.", <-osSignals)
}
