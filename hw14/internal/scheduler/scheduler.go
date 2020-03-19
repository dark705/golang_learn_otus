package sheduler

import (
	"encoding/json"
	"sync"
	"time"

	"github.com/dark705/otus/hw14/internal/config"
	"github.com/dark705/otus/hw14/internal/rabbitmq"
	"github.com/dark705/otus/hw14/internal/storage"
	"github.com/sirupsen/logrus"
)

type Scheduler struct {
	log  *logrus.Logger
	conf config.Config
	stor storage.Interface
	rmq  *rabbitmq.RMQ
	done chan struct{}
	wg   *sync.WaitGroup
}

func NewScheduler(c config.Config, l *logrus.Logger, s storage.Interface, r *rabbitmq.RMQ) *Scheduler {
	return &Scheduler{
		log:  l,
		conf: c,
		stor: s,
		rmq:  r,
		done: make(chan struct{}),
		wg:   &sync.WaitGroup{},
	}
}

func (s *Scheduler) Shutdown() {
	close(s.done)
	s.wg.Wait()
}

func (s *Scheduler) Run() {
	go func() {
		//connect to DB
		s.wg.Add(1)
		defer s.wg.Done()
		ticker := time.NewTicker(time.Second * time.Duration(s.conf.SchedulerCheckInSeconds))
		defer ticker.Stop()
		s.log.Infoln("Started Scheduler")
		for {
			select {
			case <-ticker.C:
				events, err := s.stor.GetAllNotScheduled()
				if err != nil {
					s.log.Errorln("Err on get not scheduled events from db", err)
					continue
				}
				if len(events) == 0 {
					s.log.Debugln("No notify need to be send")
					continue
				}
				for _, event := range events {
					if time.Now().Add(time.Second * time.Duration(s.conf.SchedulerNotifyInSeconds)).Before(event.StartTime) {
						s.log.Debugln("Too early send notice for event", event)
						continue
					}
					message, _ := json.Marshal(event)
					err = s.rmq.Send(message)
					if err != nil {
						s.log.Errorln("Fail send notify to RMQ:", message)
						err := s.rmq.Reconnect()
						if err != nil {
							s.log.Errorln("Fail reconnect to RMQ:", message)
						}
						continue
					}
					s.log.Infoln("Success send notify to RMQ:", string(message))
					event.IsScheduled = true
					err = s.stor.Edit(event)
					if err != nil {
						s.log.Errorln("Fail mark event in db as scheduled", err)
						continue
					}
					s.log.Infoln("Success mark event in db as scheduled", err)
				}
			case <-s.done:
				s.log.Infoln("Shutdown scheduler")
				return
			}
		}
	}()
}
