package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/dark705/otus/hw14/internal/config"
	"github.com/dark705/otus/hw14/internal/helpers"
	"github.com/dark705/otus/hw14/internal/logger"
	"github.com/dark705/otus/hw14/internal/rabbitmq"
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

	done := make(chan struct{}, 1)

	scheduler := sync.WaitGroup{}

	//DB connect
	stor, err := storage.NewPG(conf.Pg, &log)
	helpers.FailOnError(err, "postgres fail")

	//RMQ connect
	rmq, err := rabbitmq.NewRMQ(conf.Rmq, &log)
	helpers.FailOnError(err, "RMQ fail")

	//Scheduler
	go func() {
		//connect to DB
		scheduler.Add(1)
		defer scheduler.Done()
		ticker := time.NewTicker(time.Second * time.Duration(conf.SchedulerCheckInSeconds))
		for {
			select {
			case <-ticker.C:
				events, err := stor.GetAllNotScheduled()
				if err != nil {
					log.Errorln("Err on get not scheduled events from db", err)
					continue
				}
				if len(events) == 0 {
					log.Debugln("No notify need to be send")
					continue
				}
				for _, event := range events {
					if time.Now().Add(time.Second * time.Duration(conf.SchedulerNotifyInSeconds)).Before(event.StartTime) {
						log.Debugln("Too early send notice for event", event)
						continue
					}
					message, _ := json.Marshal(event)
					err = rmq.Send(message)
					if err != nil {
						log.Errorln("Fail send notify to RMQ:", message)
						_ = rmq.Reconnect()
						continue
					}
					log.Infoln("Success send notify to RMQ:", string(message))
					event.IsScheduled = true
					err = stor.Edit(event)
					if err != nil {
						log.Errorln("Fail mark event in db as scheduled", err)
						continue
					}
					log.Infoln("Success mark event in db as scheduled", err)
				}
			case <-done:
				log.Infoln("Shutdown scheduler")
				return
			}
		}
	}()

	log.Infof("Got signal from OS: %v. Exit.", <-osSignals)
	close(done)
	scheduler.Wait()
	rmq.Shutdown()
	stor.Shutdown()
}
