package calendar

type BusinessError string

func (e BusinessError) Error() string {
	return string(e)
}

const (
	ErrDateBusy          = BusinessError("Date interval already busy by another event")
	ErrNoEventsInStorage = BusinessError("No events in storage")
)
