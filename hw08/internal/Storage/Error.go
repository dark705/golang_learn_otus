package Storage

import (
	"errors"
	"fmt"
)

func ErrDateBusy() error {
	return errors.New("Date interval already busy by another event")
}

func ErrNotFoundWithId(id int) error {
	return errors.New(fmt.Sprintf("Event with id:%d not found", id))
}
