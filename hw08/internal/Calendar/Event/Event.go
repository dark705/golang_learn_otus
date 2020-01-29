package Event

import "time"

type Event struct {
	Id          int
	StartTime   time.Time
	EndTime     time.Time
	Title       string
	Description string
}

func GetEvent(startTime, endTime, title, description string) (Event, error) {
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
