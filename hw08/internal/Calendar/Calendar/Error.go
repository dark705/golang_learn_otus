package Calendar

type ErrDateBusy string

func (ErrDateBusy) Error() string {
	return "Date interval already busy by another event"
}
