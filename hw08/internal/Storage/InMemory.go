package Storage

import (
	"github.com/dark705/otus/hw08/internal/Calendar/Event"
)

type InMemory struct {
	Events map[int]Event.Event
}

func (s *InMemory) Init() {
	s.Events = make(map[int]Event.Event)
}

func (s *InMemory) Add(e Event.Event) error {
	e.Id = len(s.Events)
	if s.intervalIsBusy(e) {
		return ErrDateBusy()
	}
	s.Events[len(s.Events)] = e
	return nil
}

func (s *InMemory) Del(id int) error {
	_, exist := s.Events[id]
	if !exist {
		return ErrNotFoundWithId(id)
	}
	delete(s.Events, id)
	return nil
}

func (s *InMemory) Get(id int) (Event.Event, error) {
	event, exist := s.Events[id]
	if !exist {
		return Event.Event{}, ErrNotFoundWithId(id)
	}
	return event, nil
}

func (s *InMemory) Edit(id int, e Event.Event) error {
	_, exist := s.Events[id]
	if !exist {
		return ErrNotFoundWithId(id)
	}
	if s.intervalIsBusy(e) {
		return ErrDateBusy()
	}
	s.Events[id] = e
	return nil
}

func (s *InMemory) intervalIsBusy(newEvent Event.Event) bool {

	for id, existEvent := range s.Events {
		if newEvent.Id == id {
			continue
		}
		//NewEvent include existEvent
		if newEvent.StartTime.Before(existEvent.StartTime) && newEvent.EndTime.After(existEvent.EndTime) {
			return true
		}
		//EndTime of newEvent inside existEvent
		if newEvent.EndTime.After(existEvent.StartTime) && newEvent.EndTime.Before(existEvent.EndTime) {
			return true
		}
		//StartTime of newEvent inside existEvent
		if newEvent.StartTime.After(existEvent.StartTime) && newEvent.StartTime.Before(existEvent.EndTime) {
			return true
		}

	}
	return false
}
