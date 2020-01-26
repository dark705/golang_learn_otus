package Logger

import (
	"os"

	"github.com/dark705/otus/hw08/internal/Config"
	"github.com/sirupsen/logrus"
)

func GetLogger(c Config.Config) logrus.Logger {
	logger := logrus.Logger{}

	switch c.LogLevel {
	case "info":
		logger.SetLevel(logrus.InfoLevel)
	case "debug":
		fallthrough
	default:
		logger.SetLevel(logrus.DebugLevel)

	}

	formatter := logrus.JSONFormatter{}
	logger.SetFormatter(&formatter)
	logger.SetOutput(os.Stdout)
	return logger
}
