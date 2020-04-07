package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Api struct {
		HttpListen           string `yaml:"http_listen"`
		PrometheusHttpListen string `yaml:"prometheus_http_listen"`
		GrpcListen           string `yaml:"grpc_listen"`
		PrometheusGrpcListen string `yaml:"prometheus_grpc_listen"`
	} `yaml:"api"`
	Sender struct {
		NumOfSenders     int    `yaml:"num_of_senders"`
		PrometheusListen string `yaml:"prometheus_listen"`
	} `yaml:"sender"`
	Scheduler struct {
		CheckInSeconds  int `yaml:"check_in_seconds"`
		NotifyInSeconds int `yaml:"notify_in_seconds"`
	} `yaml:"scheduler"`
	Pg struct {
		HostPort       string `yaml:"host_port"`
		User           string `yaml:"user"`
		Pass           string `yaml:"pass"`
		Database       string `yaml:"database"`
		TimeoutConnect int    `yaml:"timeout_connect"`
		TimeoutExecute int    `yaml:"timeout_execute"`
	} `yaml:"postgres"`
	Logger struct {
		File  string `yaml:"file"`
		Level string `yaml:"level"`
	} `yaml:"log"`
	Rmq struct {
		User     string `yaml:"user"`
		Pass     string `yaml:"pass"`
		HostPort string `yaml:"host_port"`
		Timeout  int    `yaml:"timeout_connect"`
		Queue    string `yaml:"queue"`
	} `yaml:"rmq"`
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
