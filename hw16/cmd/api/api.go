package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/dark705/otus/hw16/internal/calendar/calendar"
	"github.com/dark705/otus/hw16/internal/config"
	"github.com/dark705/otus/hw16/internal/grpc"
	"github.com/dark705/otus/hw16/internal/helpers"
	"github.com/dark705/otus/hw16/internal/logger"
	"github.com/dark705/otus/hw16/internal/storage"
	"github.com/dark705/otus/hw16/internal/web"
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

	//PG
	stor, err := storage.NewPG(storage.PostgresConfig{
		HostPort:       conf.Pg.HostPort,
		User:           conf.Pg.User,
		Pass:           conf.Pg.Pass,
		Database:       conf.Pg.Database,
		TimeoutConnect: conf.Pg.TimeoutConnect,
		TimeoutExecute: conf.Pg.TimeoutExecute,
	}, &log)
	helpers.FailOnError(err, "postgres fail")

	cal := calendar.Calendar{Config: conf, Storage: &stor, Logger: &log}

	//web Server
	ws := web.NewServer(web.Config{
		HttpListen:       conf.Api.HttpListen,
		PrometheusListen: conf.Api.PrometheusHttpListen},
		&log)
	ws.RunServer()

	//gRPC Server
	grpcServer := grpc.Server{Config: grpc.Config{
		GrpcListen:       conf.Api.GrpcListen,
		PrometheusListen: conf.Api.PrometheusGrpcListen},
		Logger:   &log,
		Calendar: &cal,
	}
	go grpcServer.Run()

	log.Infof("Got signal from OS: %v. Exit.", <-osSignals)
	ws.Shutdown()
	grpcServer.Shutdown()
	stor.Shutdown()
}
