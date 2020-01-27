package Storage

import (
	"github.com/dark705/otus/hw08/internal/Calendar/Event"
)

type InMemory struct {
	Events []Event.Event
}

func (s *InMemory) Add(e Event.Event) error {
	s.Events = append(s.Events, e)
	return nil
}

func (s *InMemory) CheckIntervalIsBusy(e Event.Event) bool {
	for _, v := range s.Events {
		if e.StartTime.After(v.StartTime) && e.StartTime.Before(v.EndTime) {
			return true
		}
		if e.EndTime.Before(v.EndTime) && e.EndTime.After(v.StartTime) {
			return true
		}
	}
	return false
}
