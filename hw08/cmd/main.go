package main

import (
	"fmt"
	"os"

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

	event1, _ := Event.GetEvent("2006-01-02T15:00:00Z", "2006-01-02T16:00:00Z", "Event 1", "Some Desc1")
	event2, _ := Event.GetEvent("2006-01-02T16:00:00Z", "2006-01-02T17:00:00Z", "Event 2", "Some Desc2")
	event3, _ := Event.GetEvent("2006-01-01T16:01:00Z", "2006-01-03T18:00:00Z", "Event 3", "Some Desc3")

	calendar.AddEvent(event1)
	calendar.AddEvent(event2)
	calendar.AddEvent(event3)
	fmt.Println(inMemory)
	calendar.DelEvent(1)
	fmt.Println(inMemory)
}
