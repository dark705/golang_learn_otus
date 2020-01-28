package Storage

import (
	"errors"
	"fmt"
)

func ErrDateBusy() error {
	return errors.New("date interval already busy by another event")
}

func ErrNotFoundWithId(id int) error {
	return errors.New(fmt.Sprintf("Event with id:%d not found", id))
}
