package calendar

import (
	"sync"

	"github.com/dark705/otus/hw11/internal/calendar/event"
	"github.com/dark705/otus/hw11/internal/config"
	"github.com/dark705/otus/hw11/internal/storage"
	"github.com/sirupsen/logrus"
)

type Calendar struct {
	Config  config.Config
	mu      sync.Mutex
	Storage storage.Interface
	Logger  *logrus.Logger
}

func (c *Calendar) AddEvent(e event.Event) error {
	c.Logger.Debug("Try add to storage, Event:", e)
	c.mu.Lock()
	defer c.mu.Unlock()
	isBusy, err := c.Storage.IntervalIsBusy(e, true)
	if err != nil {
		c.Logger.Debug("Fail to check interval for Event, Error:", err)
		return err
	}
	if isBusy {
		c.Logger.Warn("Interval is busy for Event:", e)
		return ErrDateBusy
	}
	err = c.Storage.Add(e)
	if err != nil {
		c.Logger.Debug("Fail add Event to storage:", err)
		return err
	}
	c.Logger.Info("Success add to storage, Event:", e)
	return nil
}

func (c *Calendar) DelEvent(id int) error {
	c.Logger.Debug("Try del form storage Event, with Id:", id)
	err := c.Storage.Del(id)
	if err != nil {
		c.Logger.Debug("Fail del Event from storage:", err)
		return err
	}
	c.Logger.Info("Success del from storage Event, with Id:", id)
	return nil
}

func (c *Calendar) GetEvent(id int) (event.Event, error) {
	c.Logger.Debug("Try get Event form storage, with Id:", id)
	e, err := c.Storage.Get(id)
	if err != nil {
		c.Logger.Debug("Fail get Event from storage:", err)
		return e, err
	}
	c.Logger.Info("Success get from storage Event, with Id:", id)
	return e, nil
}

func (c *Calendar) GetAllEvents() ([]event.Event, error) {
	c.Logger.Debug("Try get all Events form storage")
	events, err := c.Storage.GetAll()
	if err != nil {
		c.Logger.Debug("Fail get Events from storage:", err)
		return events, err
	}
	if len(events) == 0 {
		c.Logger.Warn("No events in storage")
		return events, ErrNoEventsInStorage
	}
	c.Logger.Info("Success get from storage Events", len(events))
	return events, nil
}

func (c *Calendar) EditEvent(e event.Event) error {
	c.Logger.Debug("Try edit Event in storage")
	c.mu.Lock()
	defer c.mu.Unlock()
	isBusy, err := c.Storage.IntervalIsBusy(e, false)
	if err != nil {
		c.Logger.Debug("Fail to check interval for Event, Error:", err)
		return err
	}
	if isBusy {
		c.Logger.Warn("Interval is busy for Event:", e)
		return ErrDateBusy
	}
	err = c.Storage.Edit(e)
	if err != nil {
		c.Logger.Debug("Fail edit Event in storage:", err)
		return err
	}
	c.Logger.Info("Success edit Event in storage with Id:", e.Id)
	return nil
}
