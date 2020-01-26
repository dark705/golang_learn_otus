package Calendar

import (
	"github.com/dark705/otus/hw08/internal/Calendar/Event"
	"github.com/dark705/otus/hw08/internal/Config"
	"github.com/dark705/otus/hw08/internal/Storage"
	"github.com/sirupsen/logrus"
)

type Calendar struct {
	Config  Config.Config
	Storage Storage.Storage
	Logger  logrus.Logger
}

func (c Calendar) AddEvent(e Event.Event) int {
	c.Logger.Debug("Add Event", e)
	return c.Storage.Add(e)
}
