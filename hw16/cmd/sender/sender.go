package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/dark705/otus/hw16/internal/config"
	"github.com/dark705/otus/hw16/internal/helpers"
	"github.com/dark705/otus/hw16/internal/logger"
	"github.com/dark705/otus/hw16/internal/rabbitmq"
	"github.com/dark705/otus/hw16/internal/sender"
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

	//RMQ connect
	rmq, err := rabbitmq.NewRMQ(rabbitmq.Config{
		User:     conf.Rmq.User,
		Pass:     conf.Rmq.Pass,
		HostPort: conf.Rmq.HostPort,
		Timeout:  conf.Rmq.Timeout,
		Queue:    conf.Rmq.Queue,
	}, &log)
	helpers.FailOnError(err, "RMQ fail")

	//Senders
	senders := sender.NewSenders(sender.Config{
		NumOfSenders:     conf.Sender.NumOfSenders,
		PrometheusListen: conf.Sender.PrometheusListen,
	}, &log, rmq, sender.SendToStdout)
	senders.Run()

	log.Infof("Got signal from OS: %v. Exit.", <-osSignals)
	senders.Shutdown()
	rmq.Shutdown()
}
