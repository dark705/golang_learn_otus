package Storage

import (
	"github.com/dark705/otus/hw08/internal/Calendar/Event"
)

type InMemory struct {
	Events []Event.Event
}

func (s *InMemory) Add(e Event.Event) int {
	s.Events = append(s.Events, e)
	return len(s.Events)
}
