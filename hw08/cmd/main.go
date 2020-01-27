package main

import (
	"fmt"
	"os"
	"time"

	"github.com/dark705/otus/hw08/internal/Calendar/Calendar"
	"github.com/dark705/otus/hw08/internal/Calendar/Event"
	"github.com/dark705/otus/hw08/internal/Config"
	"github.com/dark705/otus/hw08/internal/Logger"
	"github.com/dark705/otus/hw08/internal/Storage"
)

func main() {

	config, err := Config.ReadFromFile("config/config.yaml")
	if err != nil {
		_, _ = fmt.Fprint(os.Stderr, err)
		os.Exit(2)
	}
	logger := Logger.GetLogger(config)

	inMemory := Storage.InMemory{}
	calendar := Calendar.Calendar{Config: config, Storage: &inMemory, Logger: logger}

	timeStart1, _ := time.Parse(time.RFC3339, "2006-01-02T15:00:00Z")
	timeEnd1, _ := time.Parse(time.RFC3339, "2006-01-02T16:00:00Z")
	event1 := Event.Event{timeStart1, timeEnd1, "Event 1", "Some Desc1"}

	timeStart2, _ := time.Parse(time.RFC3339, "2006-01-02T16:00:00Z")
	timeEnd2, _ := time.Parse(time.RFC3339, "2006-01-02T17:00:00Z")
	event2 := Event.Event{timeStart2, timeEnd2, "Event 2", "Some Desc2"}

	timeStart3, _ := time.Parse(time.RFC3339, "2006-01-02T16:01:00Z")
	timeEnd3, _ := time.Parse(time.RFC3339, "2006-01-02T18:00:00Z")
	event3 := Event.Event{timeStart3, timeEnd3, "Event 3", "Some Desc3"}

	calendar.AddEvent(event1)
	calendar.AddEvent(event2)
	calendar.AddEvent(event3)

	fmt.Println(inMemory)

}
