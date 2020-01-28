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
	c.Logger.Debug("Try add to storage, Event:", e)
	intervalIsBusy := c.Storage.CheckIntervalIsBusy(e)
	if intervalIsBusy {
		c.Logger.Debug("Interval is busy for Event:", e)
		return ErrDateBusy("")
	} else {
		err := c.Storage.Add(e)
		if err != nil {
			c.Logger.Debug("Fail add to storage, Event:", e)
			return err
		}
	}
	c.Logger.Info("Success add to storage, Event:", e)
	return nil
}

func (c Calendar) DelEvent(id int) error {
	c.Logger.Debug("Try del form storage Event, with Id:", id)
	err := c.Storage.Del(id)
	if err != nil {
		c.Logger.Debug("Fail del from storage Event, with Id:", id)
		return err
	}
	c.Logger.Info("Success del from storage Event, with Id:", id)
	return nil
}
