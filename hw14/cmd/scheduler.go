package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/dark705/otus/hw14/internal/config"
	"github.com/dark705/otus/hw14/internal/helpers"
	"github.com/dark705/otus/hw14/internal/logger"
	"github.com/dark705/otus/hw14/internal/rabbitmq"
	sheduler "github.com/dark705/otus/hw14/internal/scheduler"
	"github.com/dark705/otus/hw14/internal/storage"
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

	//DB connect
	stor, err := storage.NewPG(conf.Pg, &log)
	helpers.FailOnError(err, "postgres fail")

	//RMQ connect
	rmq, err := rabbitmq.NewRMQ(conf.Rmq, &log)
	helpers.FailOnError(err, "RMQ fail")

	//Scheduler
	sch := sheduler.NewScheduler(conf, &log, &stor, rmq)
	sch.Run()
	sch.Run()

	log.Infof("Got signal from OS: %v. Exit.", <-osSignals)
	sch.Shutdown()
	rmq.Shutdown()
	stor.Shutdown()
}
