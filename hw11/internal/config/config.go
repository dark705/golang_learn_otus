package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Config struct {
	HttpListen string `yaml:"http_listen"`
	GrpcListen string `yaml:"grpc_listen"`
	LogFile    string `yaml:"log_file"`
	LogLevel   string `yaml:"log_level"`
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
