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

	storage := Storage.InMemory{}
	calendar := Calendar.Calendar{Config: config, Storage: &storage, Logger: logger}

	event1 := Event.Event{time.Now(), time.Now(), "Event 1", "Some Desc1"}
	event2 := Event.Event{time.Now(), time.Now(), "Event 2", "Some Desc1"}

	calendar.AddEvent(event1)
	calendar.AddEvent(event2)

	fmt.Println(storage)

}
