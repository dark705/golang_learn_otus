package main

import (
	"errors"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/dark705/otus/hw13/internal/config"
	"github.com/dark705/otus/hw13/internal/helpers"
	"github.com/dark705/otus/hw13/internal/logger"
	"github.com/dark705/otus/hw13/internal/rabbitmq"
	"github.com/dark705/otus/hw13/internal/storage"
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

	done := make(chan struct{}, 1)

	senders := sync.WaitGroup{}

	//DB connect
	stor := storage.Postgres{Config: conf, Logger: &log}
	err = stor.Init()

	//RMQ connect
	rmq, err := rabbitmq.NewRMQ(conf.Rmq, &log)
	helpers.FailOnError(err, "RMQ fail")
	msgsCh, err := rmq.GetMsgsCh()
	helpers.FailOnError(err, "RMQ fail to get msgs chan")
	//msgsRetChan := make(chan amqp.Delivery)

	//Senders
	for i := 0; i < conf.SenderNumOfSenders; i++ {
		go func(i int) {
			senders.Add(1)
			defer senders.Done()
			for {
				select {
				case message := <-msgsCh:
					err := Send(message.Body, i)
					if err != nil {
						log.Errorln(err)
						log.Errorln("Fail send")
						message.Nack(false, true)
						log.Errorln("Return to queue")

					} else {
						log.Infoln("Success send")
						message.Ack(false)
						log.Debugln("Send Ack")
					}
				case <-done:
					return
				}
			}
		}(i)
	}

	log.Infof("Got signal from OS: %v. Exit.", <-osSignals)
	close(done)
	senders.Wait()
	rmq.Shutdown()
	stor.Shutdown()
}

func Send(msg []byte, i int) error {

	rand.Seed(time.Now().UTC().UnixNano() + int64(i))
	rnd := rand.Intn(2000) + 1000
	time.Sleep(time.Millisecond * time.Duration(rnd)) //emulate delay on sender
	fmt.Println(rnd)
	fmt.Println("Sender:", i, "SendMessage:", string(msg))

	//	if rnd < 1500 { //emulate error on send
	//		return nil
	//	}
	return errors.New("Error on sender")
}
