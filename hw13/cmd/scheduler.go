package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/dark705/otus/hw13/internal/config"
	"github.com/dark705/otus/hw13/internal/logger"
	"github.com/dark705/otus/hw13/internal/rabbitmq"
	"github.com/dark705/otus/hw13/internal/storage"
	"github.com/sirupsen/logrus"
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

	mesCh := make(chan []byte)
	done := make(chan struct{}, 1)

	wg := sync.WaitGroup{}

	//DB
	stor := storage.Postgres{Config: conf, Logger: &log}
	err = stor.Init()

	//RMQ
	rmq, err := rabbitmq.NewRMQ(conf.Rmq, &log)
	failOnError(err, "RMQ fail")

	go func(mesCh chan []byte, doneCh chan struct{}, wg sync.WaitGroup, log logrus.Logger) {
		//connect to DB
		defer stor.Shutdown()
		wg.Add(1)
		defer wg.Done()
		ticker := time.NewTicker(time.Second * 1)
		for {
			select {
			case <-ticker.C:
				events, err := stor.GetAllNotScheduled()
				if err == nil {
					if len(events) == 0 {
						log.Debugln("No events need to be send")
						continue
					}
					for _, event := range events {
						message, err := json.Marshal(event)
						if err == nil {
							mesCh <- message
						}
					}
				} else {
					log.Error("Err", err)
				}

			case <-doneCh:
				log.Infoln("Shutdown Message sender")
				return
			}

		}
	}(mesCh, done, wg, log)

	//RMQ
	go func(mesCh chan []byte, doneCh chan struct{}, wg sync.WaitGroup, log logrus.Logger) {
		wg.Add(1)
		defer wg.Done()
		for {
			select {
			case m := <-mesCh:
				err := rmq.Send(m)
				if err != nil {
					log.Errorln("Fail send to RMQ message:", m)
					//TODO Recconect?
					return
				}
				log.Infoln("Success send to RMQ message:", m)
			case <-doneCh:
				log.Infoln("Shutdown RMQ")
				return
			}

		}

	}(mesCh, done, wg, log)

	log.Infof("Got signal from OS: %v. Exit.", <-osSignals)
	close(done)
	wg.Wait()

}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
