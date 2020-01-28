package Storage

import (
	"fmt"
)

type ErrDateBusy struct{ s string }

func (e *ErrDateBusy) Error() string {
	return e.s
}

func NewErrDateBusy() error {
	return &ErrDateBusy{"Date interval already busy by another event"}
}

type ErrNotFoundWithId struct{ s string }

func (e *ErrNotFoundWithId) Error() string {
	return e.s
}

func NewErrNotFoundWithId(id int) error {
	return &ErrNotFoundWithId{fmt.Sprintf("Event with id:%d not found", id)}
}
