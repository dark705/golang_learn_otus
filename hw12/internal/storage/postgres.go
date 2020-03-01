package storage

import (
	"context"
	"errors"
	"fmt"
	"time"

	_ "github.com/jackc/pgx/stdlib"

	"github.com/dark705/otus/hw12/internal/calendar/event"
	"github.com/dark705/otus/hw12/internal/config"
	"github.com/jmoiron/sqlx"
)

type Postgres struct {
	conf    config.Config
	db      *sqlx.DB
	ctxExec context.Context
}

func (s *Postgres) Init() (err error) {
	ctxConnect, _ := context.WithTimeout(context.Background(), time.Second*time.Duration(s.conf.PgTimeoutConnect))
	s.db, err = sqlx.ConnectContext(ctxConnect, "pgx", fmt.Sprintf("postgres://%s:%s@%s/%s", s.conf.PgUser, s.conf.PgPass, s.conf.PgHostPort, s.conf.PgDatabase))
	if err != nil {
		return err
	}
	s.ctxExec, _ = context.WithCancel(context.Background())
	return nil
}

func (s *Postgres) Shutdown() error {
	return s.db.Close()
}

func (s *Postgres) Add(e event.Event) (err error) {
	sql := "INSERT INTO events (start_time, end_time, title, description) VALUES (:start_time, :end_time, :title, :description);" // :start_time from `db:"start_time"` and so on
	_, err = s.db.NamedExecContext(s.ctxExec, sql, e)
	return err
}

func (s *Postgres) Del(id int) (err error) {
	sql := "DELETE FROM events WHERE id = :id;"
	res, err := s.db.NamedExecContext(s.ctxExec, sql, struct {
		Id int `db:"id"`
	}{Id: id})
	if err != nil {
		return err
	}
	aff, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if aff != 1 {
		return errors.New(fmt.Sprintf("Fail to del, no event with %d:", id))
	}
	return nil
}

func (s *Postgres) Get(id int) (e event.Event, err error) {
	sql := "SELECT * FROM events WHERE id = :id;" // :id from `db:"id"`
	rows, err := s.db.NamedQueryContext(s.ctxExec, sql, struct {
		Id int `db:"id"`
	}{Id: id})
	if err != nil {
		return e, err
	}

	rows.Next()
	if err := rows.StructScan(&e); err != nil {
		return e, err
	}

	return e, err
}

func (s *Postgres) GetAll() (events []event.Event, err error) {
	sql := "SELECT * FROM events;" // :id from `db:"id"`
	rows, err := s.db.QueryxContext(s.ctxExec, sql)
	if err != nil {
		return events, err
	}

	for rows.Next() {
		var e event.Event
		if err = rows.StructScan(&e); err != nil {
			return events, err
		}
		events = append(events, e)
	}

	return events, err
}

func (s *Postgres) Edit(editEvent event.Event) (err error) {
	sql := "UPDATE events SET (start_time, end_time, title, description) = (:start_time, :end_time, :title, :description) WHERE id = :id;" // :start_time from `db:"start_time"` and so on
	res, err := s.db.NamedExecContext(s.ctxExec, sql, editEvent)
	if err != nil {
		return err
	}
	aff, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if aff != 1 {
		return errors.New(fmt.Sprintf("Fail to edit, no event with %d:", editEvent.Id))
	}

	return nil
}

func (s *Postgres) IntervalIsBusy(newEvent event.Event, new bool) (exist bool, err error) {
	var rows *sqlx.Rows
	//SELECT * FROM events WHERE start_time < 'new_end_time' and end_time > 'new_start_time'
	if new == true {
		rows, err = s.db.NamedQueryContext(s.ctxExec, "SELECT true FROM events WHERE start_time < :end_time AND end_time > :start_time;", newEvent)
	} else {
		rows, err = s.db.NamedQueryContext(s.ctxExec, "SELECT true FROM events WHERE start_time < :end_time AND end_time > :start_time and id != :id;", newEvent)
	}
	if err != nil {
		return exist, err
	}
	exist = rows.Next()

	return exist, err
}
