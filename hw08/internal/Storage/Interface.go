package Storage

import "github.com/dark705/otus/hw08/internal/Calendar/Event"

type Storage interface {
	Add(e Event.Event) error
	Del(id int) error
	Get(id int) (Event.Event, error)
	//Edit(id int, e Event.Event)  error
}
