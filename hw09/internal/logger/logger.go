package logger

import (
	"fmt"
	"io"
	"os"

	"github.com/dark705/otus/hw08/internal/config"
	"github.com/sirupsen/logrus"
)

var file *os.File

func GetLogger(c config.Config) logrus.Logger {
	logger := logrus.Logger{}
	switch c.LogLevel {
	case "error":
		logger.SetLevel(logrus.ErrorLevel)
	case "warn":
		logger.SetLevel(logrus.WarnLevel)
	case "info":
		logger.SetLevel(logrus.InfoLevel)
	case "debug":
		fallthrough
	default:
		logger.SetLevel(logrus.DebugLevel)

	}

	formatter := logrus.JSONFormatter{}
	logger.SetFormatter(&formatter)
	var err error
	file, err = os.OpenFile(c.LogFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		_, _ = fmt.Fprint(os.Stderr, err)
		os.Exit(2)
	}

	mw := io.MultiWriter(os.Stdout, file)
	logger.SetOutput(mw)
	return logger
}

func CloseLogFile() {
	err := file.Close()
	if err != nil {
		_, _ = fmt.Fprint(os.Stderr, err)
	}
}
