package storage

import (
	"testing"
	"time"

	"github.com/dark705/otus/hw08/internal/calendar/event"
)

func TestNewStorageHaveNoEvents(t *testing.T) {
	inMemory := InMemory{}
	inMemory.Init()
	events, err := inMemory.GetAll()
	if err == nil || len(events) != 0 {
		t.Error("In new storage exist events")
	}
}

func TestAddEventSuccess(t *testing.T) {
	inMemory := InMemory{}
	inMemory.Init()

	event, _ := event.CreateEvent("2006-01-02T15:00:00Z", "2006-01-02T16:00:00Z", "Event 1", "Some Desc1")
	err := inMemory.Add(event)
	if err != nil {
		t.Error("Can't add event to storage")
	}

	events, err := inMemory.GetAll()
	if err != nil || len(events) != 1 {
		t.Error("In storage not 1 event")
	}
}

func TestDelEventSuccess(t *testing.T) {
	inMemory := InMemory{}
	inMemory.Init()

	event, _ := event.CreateEvent("2006-01-02T15:00:00Z", "2006-01-02T16:00:00Z", "Event 1", "Some Desc1")
	_ = inMemory.Add(event)
	err := inMemory.Del(0)
	if err != nil {
		t.Error("Can't del event from storage")
	}

	events, err := inMemory.GetAll()
	if err == nil || len(events) != 0 {
		t.Error("In storage exist events")
	}
}

func TestAddDateIntervalBusy(t *testing.T) {
	var err error
	inMemory := InMemory{}
	inMemory.Init()

	event1, _ := event.CreateEvent("2006-01-02T15:00:00Z", "2006-01-02T16:00:00Z", "Event 1", "Some Desc1")
	event2, _ := event.CreateEvent("2006-01-02T16:00:00Z", "2006-01-02T17:00:00Z", "Event 2", "Some Desc2")
	event3, _ := event.CreateEvent("2006-01-02T18:00:00Z", "2006-01-02T19:00:00Z", "Event 3", "Some Desc3")
	err = inMemory.Add(event1)
	err = inMemory.Add(event2)
	err = inMemory.Add(event3)
	if err != nil {
		t.Error("Error on add not intersection events")
	}

	event4, _ := event.CreateEvent("2006-01-02T16:10:00Z", "2006-01-02T16:20:00Z", "Event 4", "Some Desc4")
	err = inMemory.Add(event4)
	if err == nil {
		t.Error("Add not return error for busy interval")
	}

	event5, _ := event.CreateEvent("2006-01-02T10:10:00Z", "2006-01-02T22:00:00Z", "Event 5", "Some Desc5")
	err = inMemory.Add(event5)
	if err == nil {
		t.Error("Add not return error for busy interval")
	}

	event6, _ := event.CreateEvent("2006-01-02T17:10:00Z", "2006-01-02T18:10:00Z", "Event 6", "Some Desc6")
	err = inMemory.Add(event6)
	if err == nil {
		t.Error("Add not return error for busy interval")
	}
}

func TestGetEvent(t *testing.T) {
	inMemory := InMemory{}
	inMemory.Init()

	event, _ := event.CreateEvent("2006-01-02T15:00:00Z", "2006-01-02T16:00:00Z", "Event 1", "Some Desc1")
	_ = inMemory.Add(event)

	_, err := inMemory.Get(1)
	if err == nil {
		t.Error("Not get error, for not exist event")
	}
	getEvent, err := inMemory.Get(0)
	if err != nil {
		t.Error("Get error, for exist event")
	}

	if getEvent.StartTime != event.StartTime ||
		getEvent.EndTime != event.EndTime ||
		getEvent.Title != event.Title ||
		getEvent.Description != event.Description {
		t.Error("Event in storage not ident")
	}
}

func TestEditEvent(t *testing.T) {
	inMemory := InMemory{}
	inMemory.Init()

	event, _ := event.CreateEvent("2006-01-02T15:00:00Z", "2006-01-02T16:00:00Z", "Event 1", "Some Desc1")
	_ = inMemory.Add(event)

	editEvent, _ := inMemory.Get(0)
	editEvent.StartTime, _ = time.Parse(time.RFC3339, "2006-01-02T15:10:00Z")
	editEvent.EndTime, _ = time.Parse(time.RFC3339, "2006-01-02T15:20:00Z")
	editEvent.Title = "newTitle"
	editEvent.Description = "newDescription"

	err := inMemory.Edit(editEvent)
	if err != nil {
		t.Error("Got not expected error on edit")
	}

	eventFromStorageAfterEdit, _ := inMemory.Get(0)
	if eventFromStorageAfterEdit != editEvent {
		t.Error("Edit Event not ident Event in storage after edit")
	}
}
