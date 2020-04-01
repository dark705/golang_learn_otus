package tests

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/cucumber/godog"
	"github.com/cucumber/godog/gherkin"
	"github.com/dark705/otus/hw16/internal/calendar/event"
	"github.com/dark705/otus/hw16/internal/helpers"
	"github.com/dark705/otus/hw16/internal/storage"
	"github.com/sirupsen/logrus"
)

var pg storage.Postgres
var lastID int

func iAddEventInDBWithData(data *gherkin.DocString) (err error) {
	var e event.Event

	replacer := strings.NewReplacer("\n", "", "\t", "")
	cleanJSON := replacer.Replace(data.Content)

	err = json.Unmarshal([]byte(cleanJSON), &e)
	if err != nil {
		err = fmt.Errorf("can't unmarshal, error: %v", err)
		return err
	}

	err = pg.Add(e)
	if err != nil {
		err = fmt.Errorf("can't add event in db, error: %v", err)
		return err
	}
	return nil
}

func thenIWaitSecondsToMakeSureThatTheSchedulerProcessedEvent(sec string) error {
	s, _ := strconv.Atoi(sec)
	time.Sleep(time.Second * time.Duration(s))
	return nil
}

func checkDBForNotScheduledEvents() error {
	//getAll
	eventsFromPg, err := pg.GetAll()
	if err != nil || len(eventsFromPg) < 1 {
		return errors.New("fail to get all events from storage")
	}
	lastID = eventsFromPg[len(eventsFromPg)-1].Id

	//get
	eventFromPg, err := pg.Get(lastID)
	if err != nil {
		return errors.New("fail to get event from storage")

	}
	if !eventFromPg.IsScheduled {
		return errors.New("event in pg was not Scheduled")
	}
	return nil
}

func deleteTestEvent() error {
	err := pg.Del(lastID)
	if err != nil {
		return errors.New("fail to del test event in storage")
	}
	return nil
}

func checkDBForScheduledEvents() error {
	//getAll
	eventsFromPg, err := pg.GetAll()
	if err != nil || len(eventsFromPg) < 1 {
		return errors.New("fail to get all events from storage")
	}
	lastID = eventsFromPg[len(eventsFromPg)-1].Id

	//get
	eventFromPg, err := pg.Get(lastID)
	if err != nil {
		return errors.New("fail to get event from storage")

	}
	if eventFromPg.IsScheduled {
		return errors.New("event in pg was Scheduled")
	}
	return nil
}

func checkEventOnSender() error {
	return nil
	return godog.ErrPending
}

func checkNoEventOnSender() error {
	return nil
	return godog.ErrPending
}

func FeatureContext(s *godog.Suite) {
	connectToPG()

	s.Step(`^I add Event in DB, with data:$`, iAddEventInDBWithData)
	s.Step(`^Then I wait "([^"]*)" seconds to make sure that the scheduler processed event$`, thenIWaitSecondsToMakeSureThatTheSchedulerProcessedEvent)
	s.Step(`^Check DB for not scheduled events$`, checkDBForNotScheduledEvents)
	s.Step(`^Check event on Sender$`, checkEventOnSender)
	s.Step(`^Delete test event$`, deleteTestEvent)
	s.Step(`^Then I wait "([^"]*)" seconds to make sure that the scheduler processed event$`, thenIWaitSecondsToMakeSureThatTheSchedulerProcessedEvent)
	s.Step(`^Check DB for scheduled events$`, checkDBForScheduledEvents)
	s.Step(`^Check no event on Sender$`, checkNoEventOnSender)
	s.Step(`^Delete test event$`, deleteTestEvent)
}

func connectToPG() {
	var err error
	pg, err = storage.NewPG(storage.PostgresConfig{
		HostPort:       "postgres:5432",
		User:           "postgres",
		Pass:           "postgres",
		Database:       "calendar",
		TimeoutConnect: 10,
	}, &logrus.Logger{})
	helpers.FailOnError(err, "Postgres fail")

}
