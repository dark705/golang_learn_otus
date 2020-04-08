package storage

import (
	"errors"
	"fmt"

	"github.com/dark705/otus/hw17/internal/calendar/event"
)

type InMemory struct {
	Events map[int]event.Event
	LastId int
}

func (s *InMemory) Init() error {
	s.Events = make(map[int]event.Event)
	return nil
}

func (s *InMemory) Add(e event.Event) error {
	e.Id = len(s.Events) + 1
	s.Events[len(s.Events)+1] = e
	return nil
}

func (s *InMemory) Del(id int) error {
	delete(s.Events, id)
	return nil
}

func (s *InMemory) Get(id int) (event.Event, error) {
	ev, exist := s.Events[id]
	if !exist {
		return ev, errors.New(fmt.Sprintf("Event with id: %d not found", id))
	}
	return ev, nil
}

func (s *InMemory) GetAll() ([]event.Event, error) {
	events := make([]event.Event, 0, len(s.Events))
	for _, e := range s.Events {
		events = append(events, e)
	}
	return events, nil
}

func (s *InMemory) GetAllNotScheduled() ([]event.Event, error) {
	events := make([]event.Event, 0, len(s.Events))
	for _, e := range s.Events {
		if !e.IsScheduled {
			events = append(events, e)
		}
	}
	return events, nil
}

func (s *InMemory) Edit(e event.Event) error {
	_, exist := s.Events[e.Id]
	if !exist {
		return errors.New(fmt.Sprintf("Event with id: %d not found", e.Id))
	}
	s.Events[e.Id] = e
	return nil
}

func (s *InMemory) IntervalIsBusy(checkedEvent event.Event, new bool) (bool, error) {
	for _, existEvent := range s.Events {
		if existEvent.StartTime.Before(checkedEvent.EndTime) && existEvent.EndTime.After(checkedEvent.StartTime) && checkedEvent.Id != existEvent.Id {
			return true, nil
		}
	}
	return false, nil
}
