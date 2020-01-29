package Storage

import "github.com/dark705/otus/hw08/internal/Calendar/Event"

type Interface interface {
	Add(e Event.Event) error
	Del(id int) error
	Get(id int) (Event.Event, error)
	GetAll() ([]Event.Event, error)
	Edit(Event.Event) error
}
