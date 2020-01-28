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
	err := c.Storage.Add(e)
	if err != nil {
		c.Logger.Debug("Fail add to storage, error:", err)
		return err
	}
	c.Logger.Debug("Success add event to storage, event:", e)
	return nil
}

func (c Calendar) DelEvent(id int) error {
	c.Logger.Debug("Try del form storage Event, with Id:", id)
	err := c.Storage.Del(id)
	if err != nil {
		c.Logger.Debug("Fail del from storage Event", " Error:", err)
		return err
	}
	c.Logger.Info("Success del from storage Event, with Id:", id)
	return nil
}

func (c Calendar) GetEvent(id int) (Event.Event, error) {
	c.Logger.Debug("Try get Event form storage, with Id:", id)
	event, err := c.Storage.Get(id)
	if err != nil {
		c.Logger.Debug("Fail get Event from storage", " Error:", err)
		return event, err
	}
	c.Logger.Info("Success get from storage Event, with Id:", id)
	return event, nil
}
