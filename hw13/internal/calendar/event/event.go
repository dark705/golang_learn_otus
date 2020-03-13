package event

import "time"

type Event struct {
	Id          int       `db:"id" json:"id"`
	StartTime   time.Time `db:"start_time" json:"startTime"`
	EndTime     time.Time `db:"end_time" json:"endTime"`
	Title       string    `db:"title" json:"title"`
	Description string    `db:"description" json:"description"`
	IsScheduled bool      `db:"is_scheduled" json:"isScheduled"`
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
