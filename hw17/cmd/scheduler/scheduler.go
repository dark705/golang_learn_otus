package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/dark705/otus/hw17/internal/config"
	"github.com/dark705/otus/hw17/internal/helpers"
	"github.com/dark705/otus/hw17/internal/logger"
	"github.com/dark705/otus/hw17/internal/rabbitmq"
	sheduler "github.com/dark705/otus/hw17/internal/scheduler"
	"github.com/dark705/otus/hw17/internal/storage"
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

	log := logger.NewLogger(logger.Config{
		File:  conf.Logger.File,
		Level: conf.Logger.Level,
	})
	defer logger.CloseLogFile()

	osSignals := make(chan os.Signal, 1)
	signal.Notify(osSignals, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)

	//DB connect
	stor, err := storage.NewPG(storage.PostgresConfig{
		HostPort:       conf.Pg.HostPort,
		User:           conf.Pg.User,
		Pass:           conf.Pg.Pass,
		Database:       conf.Pg.Database,
		TimeoutConnect: conf.Pg.TimeoutConnect,
		TimeoutExecute: conf.Pg.TimeoutExecute,
	}, &log)
	helpers.FailOnError(err, "postgres fail")

	//RMQ connect
	rmq, err := rabbitmq.NewRMQ(rabbitmq.Config{
		User:     conf.Rmq.User,
		Pass:     conf.Rmq.Pass,
		HostPort: conf.Rmq.HostPort,
		Timeout:  conf.Rmq.Timeout,
		Queue:    conf.Rmq.Queue,
	}, &log)
	helpers.FailOnError(err, "RMQ fail")

	//Scheduler
	sch := sheduler.NewScheduler(sheduler.Config{
		CheckInSeconds:  conf.Scheduler.CheckInSeconds,
		NotifyInSeconds: conf.Scheduler.NotifyInSeconds,
	}, &log, &stor, rmq)
	sch.Run()

	log.Infof("Got signal from OS: %v. Exit.", <-osSignals)
	sch.Shutdown()
	rmq.Shutdown()
	stor.Shutdown()
}
