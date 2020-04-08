package calendar

import (
	"testing"
	"time"

	"github.com/dark705/otus/hw17/internal/calendar/event"
	"github.com/dark705/otus/hw17/internal/storage"
	"github.com/sirupsen/logrus"
)

func TestNewCalendarHaveNoEvents(t *testing.T) {
	inMemory := storage.InMemory{}
	err := inMemory.Init()
	if err != nil {
		t.Error("Can't init storage")
	}
	calendar := Calendar{Storage: &inMemory, Logger: &logrus.Logger{}}

	events, err := calendar.GetAllEvents()
	if err != ErrNoEventsInStorage || len(events) != 0 {
		t.Error("In new storage exist events")
	}
}

func TestAddEventSuccess(t *testing.T) {
	inMemory := storage.InMemory{}
	err := inMemory.Init()
	if err != nil {
		t.Error("Can't init storage")
	}
	calendar := Calendar{Storage: &inMemory, Logger: &logrus.Logger{}}

	event1, _ := event.CreateEvent("2006-01-02T15:00:00+03:00", "2006-01-02T16:00:00+03:00", "Event 1", "Some Desc1")
	err = calendar.AddEvent(event1)
	if err != nil {
		t.Error("Can't add event to storage")
	}

	events, err := calendar.GetAllEvents()
	if err != nil || len(events) != 1 {
		t.Error("In storage not 1 event")
	}
}

func TestDelEventSuccess(t *testing.T) {
	inMemory := storage.InMemory{}
	err := inMemory.Init()
	if err != nil {
		t.Error("Can't init storage")
	}
	calendar := Calendar{Storage: &inMemory, Logger: &logrus.Logger{}}

	event1, _ := event.CreateEvent("2006-01-02T15:00:00+03:00", "2006-01-02T16:00:00+03:00", "Event 1", "Some Desc1")
	_ = calendar.AddEvent(event1)
	err = calendar.DelEvent(1)
	if err != nil {
		t.Error("Can't del event from storage")
	}

	events, err := calendar.GetAllEvents()
	if err == nil || len(events) != 0 {
		t.Error("In storage exist events")
	}
}

func TestAddDateIntervalBusyAtSameTime(t *testing.T) {
	var err error
	inMemory := storage.InMemory{}
	err = inMemory.Init()
	if err != nil {
		t.Error("Can't init storage")
	}
	calendar := Calendar{Storage: &inMemory, Logger: &logrus.Logger{}}

	event1, _ := event.CreateEvent("2006-01-02T15:00:00+03:00", "2006-01-02T16:00:00+03:00", "Event 1", "Some Desc1")
	event2, _ := event.CreateEvent("2006-01-02T15:00:00+03:00", "2006-01-02T16:00:00+03:00", "Event 2", "Some Desc2")
	err = calendar.AddEvent(event1)
	err = calendar.AddEvent(event2)

	if err != ErrDateBusy {
		t.Error("Add not return error for busy interval")
	}
}

func TestAddDateIntervalStartTimeInExistInterval(t *testing.T) {
	var err error
	inMemory := storage.InMemory{}
	err = inMemory.Init()
	if err != nil {
		t.Error("Can't init storage")
	}
	calendar := Calendar{Storage: &inMemory, Logger: &logrus.Logger{}}

	event1, _ := event.CreateEvent("2006-01-02T15:00:00+03:00", "2006-01-02T16:00:00+03:00", "Event 1", "Some Desc1")
	event2, _ := event.CreateEvent("2006-01-02T15:30:00+03:00", "2006-01-02T17:00:00+03:00", "Event 2", "Some Desc2")
	err = calendar.AddEvent(event1)
	err = calendar.AddEvent(event2)

	if err != ErrDateBusy {
		t.Error("Add not return error for busy interval")
	}
}

func TestAddDateIntervalEndTimeInExistInterval(t *testing.T) {
	var err error
	inMemory := storage.InMemory{}
	err = inMemory.Init()
	if err != nil {
		t.Error("Can't init storage")
	}
	calendar := Calendar{Storage: &inMemory, Logger: &logrus.Logger{}}

	event1, _ := event.CreateEvent("2006-01-02T15:00:00+03:00", "2006-01-02T16:00:00+03:00", "Event 1", "Some Desc1")
	event2, _ := event.CreateEvent("2006-01-02T14:00:00+03:00", "2006-01-02T15:30:00+03:00", "Event 2", "Some Desc2")
	err = calendar.AddEvent(event1)
	err = calendar.AddEvent(event2)

	if err != ErrDateBusy {
		t.Error("Add not return error for busy interval")
	}
}

func TestAddDateIntervalInsideExistInterval(t *testing.T) {
	var err error
	inMemory := storage.InMemory{}
	err = inMemory.Init()
	if err != nil {
		t.Error("Can't init storage")
	}
	calendar := Calendar{Storage: &inMemory, Logger: &logrus.Logger{}}

	event1, _ := event.CreateEvent("2006-01-02T15:00:00+03:00", "2006-01-02T16:00:00+03:00", "Event 1", "Some Desc1")
	event2, _ := event.CreateEvent("2006-01-02T15:10:00+03:00", "2006-01-02T15:50:00+03:00", "Event 2", "Some Desc2")
	err = calendar.AddEvent(event1)
	err = calendar.AddEvent(event2)

	if err != ErrDateBusy {
		t.Error("Add not return error for busy interval")
	}
}

func TestAddDateIntervalIncludeExistInterval(t *testing.T) {
	var err error
	inMemory := storage.InMemory{}
	err = inMemory.Init()
	if err != nil {
		t.Error("Can't init storage")
	}
	calendar := Calendar{Storage: &inMemory, Logger: &logrus.Logger{}}

	event1, _ := event.CreateEvent("2006-01-02T15:00:00+03:00", "2006-01-02T16:00:00+03:00", "Event 1", "Some Desc1")
	event2, _ := event.CreateEvent("2006-01-02T14:00:00+03:00", "2006-01-02T17:00:00+03:00", "Event 2", "Some Desc2")
	err = calendar.AddEvent(event1)
	err = calendar.AddEvent(event2)

	if err != ErrDateBusy {
		t.Error("Add not return error for busy interval")
	}
}

func TestAddDateIntervalBusyMultiple(t *testing.T) {
	var err error
	inMemory := storage.InMemory{}
	err = inMemory.Init()
	if err != nil {
		t.Error("Can't init storage")
	}
	calendar := Calendar{Storage: &inMemory, Logger: &logrus.Logger{}}

	event1, _ := event.CreateEvent("2006-01-02T15:00:00+03:00", "2006-01-02T16:00:00+03:00", "Event 1", "Some Desc1")
	event2, _ := event.CreateEvent("2006-01-02T16:00:00+03:00", "2006-01-02T17:00:00+03:00", "Event 2", "Some Desc2")
	event3, _ := event.CreateEvent("2006-01-02T18:00:00+03:00", "2006-01-02T19:00:00+03:00", "Event 3", "Some Desc3")
	err = calendar.AddEvent(event1)
	err = calendar.AddEvent(event2)
	err = calendar.AddEvent(event3)
	if err != nil {
		t.Error("Error on add not intersection events")
	}

	event4, _ := event.CreateEvent("2006-01-02T16:10:00+03:00", "2006-01-02T16:20:00+03:00", "Event 4", "Some Desc4")
	err = calendar.AddEvent(event4)
	if err != ErrDateBusy {
		t.Error("Add not return error for busy interval")
	}

	event5, _ := event.CreateEvent("2006-01-02T10:10:00+03:00", "2006-01-02T22:00:00+03:00", "Event 5", "Some Desc5")
	err = calendar.AddEvent(event5)
	if err != ErrDateBusy {
		t.Error("Add not return error for busy interval")
	}

	event6, _ := event.CreateEvent("2006-01-02T17:10:00+03:00", "2006-01-02T18:10:00+03:00", "Event 6", "Some Desc6")
	err = calendar.AddEvent(event6)
	if err != ErrDateBusy {
		t.Error("Add not return error for busy interval")
	}
}

func TestGetEvent(t *testing.T) {
	inMemory := storage.InMemory{}
	err := inMemory.Init()
	if err != nil {
		t.Error("Can't init storage")
	}
	calendar := Calendar{Storage: &inMemory, Logger: &logrus.Logger{}}

	event1, _ := event.CreateEvent("2006-01-02T15:00:00+03:00", "2006-01-02T16:00:00+03:00", "Event 1", "Some Desc1")
	event1.Id = 1
	_ = calendar.AddEvent(event1)

	_, err = calendar.GetEvent(100)

	if err == nil {
		t.Error("Not get error, for not exist event")
	}
	getEvent, err := calendar.GetEvent(1)
	if err != nil {
		t.Error("Get error, for exist event")
	}

	if getEvent != event1 {
		t.Error("Event in storage not ident")
	}
}

func TestEditEvent(t *testing.T) {
	inMemory := storage.InMemory{}
	err := inMemory.Init()
	if err != nil {
		t.Error("Can't init storage")
	}
	calendar := Calendar{Storage: &inMemory, Logger: &logrus.Logger{}}

	event1, _ := event.CreateEvent("2006-01-02T15:00:00+03:00", "2006-01-02T16:00:00+03:00", "Event 1", "Some Desc1")
	err = calendar.AddEvent(event1)
	if err != nil {
		t.Error("Got not expected error on add event")
	}

	editEvent, _ := calendar.GetEvent(1)
	editEvent.StartTime, _ = time.Parse(time.RFC3339, "2006-01-02T15:10:00+03:00")
	editEvent.EndTime, _ = time.Parse(time.RFC3339, "2006-01-02T15:20:00+03:00")
	editEvent.Title = "newTitle"
	editEvent.Description = "newDescription"

	err = calendar.EditEvent(editEvent)
	if err != nil {
		t.Error("Got not expected error on edit", err)
	}

	eventFromStorageAfterEdit, _ := calendar.GetEvent(1)
	if eventFromStorageAfterEdit != editEvent {
		t.Error("Edit Event not ident Event in storage after edit")
	}
}
