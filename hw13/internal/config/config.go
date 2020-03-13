package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Config struct {
	HttpListen       string `yaml:"http_listen"`
	GrpcListen       string `yaml:"grpc_listen"`
	LogFile          string `yaml:"log_file"`
	LogLevel         string `yaml:"log_level"`
	PgHostPort       string `yaml:"db_host_port"`
	PgUser           string `yaml:"db_user"`
	PgPass           string `yaml:"db_pass"`
	PgDatabase       string `yaml:"db_database"`
	PgTimeoutConnect int    `yaml:"db_timeout_connect"`
	PgTimeoutExecute int    `yaml:"db_timeout_execute"`
	Rmq              struct {
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
