package config

import (
	"io/ioutil"

	"github.com/dark705/otus/hw14/internal/logger"
	"github.com/dark705/otus/hw14/internal/rabbitmq"
	"github.com/dark705/otus/hw14/internal/storage"
	"gopkg.in/yaml.v2"
)

type Config struct {
	HttpListen               string                 `yaml:"http_listen"`
	GrpcListen               string                 `yaml:"grpc_listen"`
	SchedulerCheckInSeconds  int                    `yaml:"scheduler_check_in_seconds"`
	SchedulerNotifyInSeconds int                    `yaml:"scheduler_notify_in_seconds"`
	SenderNumOfSenders       int                    `yaml:"sender_num_of_senders"`
	Pg                       storage.PostgresConfig `yaml:"postgres"`
	Logger                   logger.Config          `yaml:"log"`
	Rmq                      rabbitmq.Config        `yaml:"rmq"`
}

func ReadFromFile(file string) (Config, error) {
	c := Config{}
	r, err := ioutil.ReadFile(file)
	if err != nil {
		return c, err
	}

	err = yaml.Unmarshal(r, &c)
	if err != nil {
		return c, err
	}

	return c, nil
}
