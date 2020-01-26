package Event

import "time"

type Event struct {
	StartTime   time.Time
	EndTime     time.Time
	Title       string
	Description string
}
