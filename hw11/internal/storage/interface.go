package storage

import (
	"github.com/dark705/otus/hw11/internal/calendar/event"
)

type Interface interface {
	Add(e event.Event) error
	Del(id int) error
	Get(id int) (event.Event, error)
	GetAll() ([]event.Event, error)
	Edit(event.Event) error
	IntervalIsBusy(event.Event, bool) (bool, error)
}
