package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/dark705/otus/hw14/internal/calendar/calendar"
	"github.com/dark705/otus/hw14/internal/config"
	"github.com/dark705/otus/hw14/internal/grpc"
	"github.com/dark705/otus/hw14/internal/helpers"
	"github.com/dark705/otus/hw14/internal/logger"
	"github.com/dark705/otus/hw14/internal/storage"
	"github.com/dark705/otus/hw14/internal/web"
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

	log := logger.NewLogger(conf.Logger)
	defer logger.CloseLogFile()

	osSignals := make(chan os.Signal, 1)
	signal.Notify(osSignals, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)

	//PG
	stor, err := storage.NewPG(conf.Pg, &log)
	helpers.FailOnError(err, "postgres fail")

	cal := calendar.Calendar{Config: conf, Storage: &stor, Logger: &log}

	//web Server
	ws := web.NewServer(conf, &log)
	ws.RunServer()

	//gRPC Server
	grpcServer := grpc.Server{Config: conf, Logger: &log, Calendar: &cal}
	go grpcServer.Run()

	log.Infof("Got signal from OS: %v. Exit.", <-osSignals)
	ws.Shutdown()
	grpcServer.Shutdown()
	stor.Shutdown()
}
