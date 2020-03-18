package storage

//Integration test. Work with real test PG DB only!

import (
	"testing"

	"github.com/dark705/otus/hw14/internal/calendar/event"

	"github.com/dark705/otus/hw14/internal/helpers"
	"github.com/sirupsen/logrus"
)

var conf = PostgresConfig{
	HostPort:       "127.0.0.1:54441",
	User:           "postgres",
	Pass:           "postgres",
	Database:       "calendar",
	TimeoutConnect: 10,
}

func TestAddGetAllGetDel(t *testing.T) {
	pg, err := NewPG(conf, &logrus.Logger{})
	helpers.FailOnError(err, "Postgres fail")
	defer pg.Shutdown()

	event, _ := event.CreateEvent("2006-01-02T15:00:00+03:00", "2006-01-02T15:00:00+03:00", "Event 1", "Some Desc1")

	//add
	err = pg.Add(event)
	if err != nil {
		t.Error("Fail to add event to storage")
	}

	//getAll
	eventsFromPg, err := pg.GetAll()
	if err != nil || len(eventsFromPg) < 1 {
		t.Error("Fail to get all events in storage")
	}
	lastId := eventsFromPg[len(eventsFromPg)-1].Id

	//get
	eventFromPg, err := pg.Get(lastId)
	if err != nil {
		t.Error("Fail to add event to storage")
	}
	event.Id = eventFromPg.Id //add original event id = lasId in PG for simple compare

	if eventFromPg != event {
		t.Error("Events from PG and original not same")
	}

	//del
	err = pg.Del(lastId)
	if err != nil {
		t.Error("Fail to del test event in storage")
	}
}

func TestIntervalIsBusy(t *testing.T) {
	pg, err := NewPG(conf, &logrus.Logger{})
	helpers.FailOnError(err, "Postgres fail")
	defer pg.Shutdown()

	event1, _ := event.CreateEvent("2006-01-02T15:00:00+03:00", "2006-01-02T16:00:00+03:00", "Event 1", "Some Desc1")
	event2, _ := event.CreateEvent("2006-01-02T16:00:00+03:00", "2006-01-02T17:00:00+03:00", "Event 2", "Some Desc2")
	eventIntersect, _ := event.CreateEvent("2006-01-02T15:30:00+03:00", "2006-01-02T19:00:00+03:00", "Intersect", "Intersect")

	//add1
	err = pg.Add(event1)
	if err != nil {
		t.Error("Fail to add event to storage")
	}

	//add1
	err = pg.Add(event2)
	if err != nil {
		t.Error("Fail to add event to storage")
	}

	//getAll, only for lastId event
	eventsFromPg, err := pg.GetAll()
	if err != nil || len(eventsFromPg) < 1 {
		t.Error("Fail to get all events in storage")
	}
	lastId := eventsFromPg[len(eventsFromPg)-1].Id

	//check intersect for new event
	isBusy, err := pg.IntervalIsBusy(eventIntersect, true)
	if err != nil {
		t.Error("Fail to check intersect event in storage")
	}
	if isBusy != true {
		t.Error("Interval is not busy for intersect event")
	}

	//del2
	err = pg.Del(lastId)
	if err != nil {
		t.Error("Fail to del test event2 in storage")
	}

	//del1
	err = pg.Del(lastId - 1)
	if err != nil {
		t.Error("Fail to del test event1 in storage")
	}
}

func TestEditEvent(t *testing.T) {
	pg, err := NewPG(conf, &logrus.Logger{})
	helpers.FailOnError(err, "Postgres fail")
	defer pg.Shutdown()

	event1, _ := event.CreateEvent("2006-01-02T15:00:00+03:00", "2006-01-02T16:00:00+03:00", "Event 1", "Some Desc1")
	eventEdited, _ := event.CreateEvent("2006-01-02T15:00:00+03:00", "2006-01-02T16:00:00+03:00", "EditedEvent1", "EditedEvent1")

	//add
	err = pg.Add(event1)
	if err != nil {
		t.Error("Fail to add event to PG")
	}

	//getAll, only for lastId event
	eventsFromPg, err := pg.GetAll()
	if err != nil || len(eventsFromPg) < 1 {
		t.Error("Fail to get all events in PG")
	}
	lastId := eventsFromPg[len(eventsFromPg)-1].Id

	//edit
	eventEdited.Id = lastId
	err = pg.Edit(eventEdited)
	if err != nil {
		t.Error("Fail edit event in PG")
	}

	eventFromPg, err := pg.Get(lastId)
	if err != nil {
		t.Error("Fail get event from PG")
	}

	if eventFromPg != eventEdited {
		t.Error("Events from PG and edited not same")
	}

	//del
	err = pg.Del(lastId)
	if err != nil {
		t.Error("Fail to del test event in PG")
	}
}
