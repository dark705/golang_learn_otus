package logger

import (
	"fmt"
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

type Config struct {
	File  string
	Level string
}

var file *os.File

func NewLogger(c Config) logrus.Logger {
	logger := logrus.Logger{}
	switch c.Level {
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
	file, err = os.OpenFile(c.File, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
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
