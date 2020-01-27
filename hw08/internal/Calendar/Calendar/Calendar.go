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

func (c Calendar) AddEvent(e Event.Event) error {
	c.Logger.Debug("Try add Event to storage", e)
	intervalIsBusy := c.Storage.CheckIntervalIsBusy(e)
	if intervalIsBusy {
		c.Logger.Debug("Interval for Event is busy", e)
		return ErrDateBusy("")
	} else {
		err := c.Storage.Add(e)
		if err != nil {
			c.Logger.Debug("Fail to add Event in storage", e)
			return err
		}
	}
	c.Logger.Info("Event add in storage success", e)
	return nil
}
