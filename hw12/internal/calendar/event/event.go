package event

import "time"

type Event struct {
	Id          int       `db:"id"`
	StartTime   time.Time `db:"start_time"`
	EndTime     time.Time `db:"end_time"`
	Title       string    `db:"title"`
	Description string    `db:"description"`
}

func CreateEvent(startTime, endTime, title, description string) (Event, error) {
	e := Event{}

	sTime, err := time.Parse(time.RFC3339, startTime)
	if err != nil {
		return e, err
	}
	e.StartTime = sTime

	eTime, err := time.Parse(time.RFC3339, endTime)
	if err != nil {
		return e, err
	}
	e.EndTime = eTime

	e.Title = title
	e.Description = description

	return e, nil
}
