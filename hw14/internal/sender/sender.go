package sender

import (
	"sync"

	"github.com/dark705/otus/hw14/internal/helpers"
	"github.com/dark705/otus/hw14/internal/rabbitmq"
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

type Config struct {
	NumOfSenders int `yaml:"num_of_senders"`
}

type Senders struct {
	log    *logrus.Logger
	conf   Config
	rmq    *rabbitmq.RMQ
	msgsCh <-chan amqp.Delivery
	done   chan struct{}
	wg     *sync.WaitGroup
	handle func(msg []byte, i int) error
}

func NewSenders(c Config, l *logrus.Logger, r *rabbitmq.RMQ, h func(msg []byte, i int) error) *Senders {
	msgsCh, err := r.GetMsgsCh()
	helpers.FailOnError(err, "RMQ fail to get msgs chan")

	return &Senders{
		log:    l,
		conf:   c,
		rmq:    r,
		msgsCh: msgsCh,
		done:   make(chan struct{}),
		wg:     &sync.WaitGroup{},
		handle: h,
	}
}
func (s *Senders) Shutdown() {
	s.log.Infoln("Shutdown senders")
	close(s.done)
	s.wg.Wait()
	s.log.Infoln("Success shutdown all senders")
}

func (s *Senders) Run() {
	s.log.Infoln("Run senders")
	for i := 0; i < s.conf.NumOfSenders; i++ {
		go func(i int) {
			s.wg.Add(1)
			defer func() {
				s.log.Infoln("Sender:", i, "shutdown")
				s.wg.Done()
			}()
			s.log.Infoln("Sender:", i, "waiting for notices")
			for {
				select {
				case <-s.done:
					return
				default:
					select {
					case message, ok := <-s.msgsCh:
						if !ok {
							err := s.rmq.Reconnect()
							if err != nil {
								s.log.Errorln("Fail on reconnect to RMQ", err)
							}
							continue
						}
						err := s.handle(message.Body, i)
						if err != nil {
							s.log.Errorln(err)
							s.log.Infoln("Senders:", i, "fail send")
							err := message.Nack(false, true)
							if err != nil {
								s.log.Errorln("Fail on Nack message RMQ", err)
							} else {
								s.log.Debugln("Send Nack")
							}
						} else {
							s.log.Infoln("Senders:", i, "success send")
							err := message.Ack(false)
							if err != nil {
								s.log.Errorln("Fail on Ack message RMQ", err)
							} else {
								s.log.Debugln("Send Ack")
							}

						}
					case <-s.done:
						return
					}
				}
			}
		}(i)
	}
}
