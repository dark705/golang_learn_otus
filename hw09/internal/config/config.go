package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Config struct {
	HttpListen string `yaml:"http_listen"`
	LogFile    string `yaml:"log_file"`
	LogLevel   string `yaml:"log_level"`
}

func ReadFromFile(file string) (Config, error) {
	c := Config{}
	r, e := ioutil.ReadFile(file)
	if e != nil {
		return c, e
	}

	e = yaml.Unmarshal(r, &c)
	if e != nil {
		return c, e
	}

	return c, nil
}
