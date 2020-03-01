package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/dark705/otus/hw12/internal/calendar/event"
	_ "github.com/jackc/pgx/stdlib"

	"github.com/dark705/otus/hw12/internal/config"
	"github.com/dark705/otus/hw12/internal/logger"
	"github.com/jmoiron/sqlx"
)

func main() {
	var cFile string
	flag.StringVar(&cFile, "config", "config/config.yaml", "Config file")
	flag.Parse()
	if cFile == "" {
		_, _ = fmt.Fprint(os.Stderr, "Not set config file")
		os.Exit(2)
	}

	conf, err := config.ReadFromFile(cFile)
	if err != nil {
		_, _ = fmt.Fprint(os.Stderr, err)
		os.Exit(2)
	}

	log := logger.GetLogger(conf)
	defer logger.CloseLogFile()
	_ = log

	ctxConnect, _ := context.WithTimeout(context.Background(), time.Second*time.Duration(conf.PgTimeoutConnect))
	db, err := sqlx.ConnectContext(ctxConnect, "pgx", fmt.Sprintf("postgres://%s:%s@%s/%s", conf.PgUser, conf.PgPass, conf.PgHostPort, conf.PgDatabase))
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()
	ctxExec, _ := context.WithTimeout(context.Background(), time.Second*time.Duration(conf.PgTimeoutExecute))
	/*
		Add(e event.Event) error
		Get(id int) (event.Event, error)
		Del(id int) error
		GetAll() ([]event.Event, error)
		Edit(event.Event) error
		IntervalIsBusy(event.Event, bool) (bool, error)
	*/

	/*
		//Add
		sql := "INSERT INTO events (start_time, end_time, title, description) VALUES (:start_time, :end_time, :title, :description);" // :start_time from `db:"start_time"` and so on
		_, err = db.NamedExecContext(ctxExec, sql, ev)
		if err != nil {
			log.Errorln("Fail on add event to PG", err)
		}

		//Get
		sql = "SELECT * FROM events WHERE id = :id;" // :id from `db:"id"`
		rows, err := db.NamedQueryContext(ctxExec, sql, struct {
			Id int `db:"id"`
		}{Id: 58})
		if err != nil {
			log.Errorln("Fail on get event from PG", err)
		}
		rows.Next()
		var e event.Event
		if err := rows.StructScan(&e); err != nil {
			log.Errorln("Fail on get event from PG", err)
		}
		fmt.Println(e)

		//Del
		sql = "DELETE FROM events WHERE id = :id;"
		res, err1 := db.NamedExecContext(ctxExec, sql, struct {
			Id int `db:"id"`
		}{Id: 65})

		if err1 != nil {
			log.Errorln("Fail on del event in PG", err)
		}
		aff, err := res.RowsAffected()
		fmt.Println(aff, err)
		if aff != 1 || err != nil {
			log.Errorln("Fail on del event in PG")
		}

		//GetAll()
		sql = "SELECT * FROM events;" // :id from `db:"id"`
		rows, err = db.QueryxContext(ctxExec, sql)
		if err != nil {
			log.Errorln("Fail on get event from PG", err)
		}
		var events []event.Event
		for rows.Next() {
			var e event.Event
			if err := rows.StructScan(&e); err != nil {
				log.Errorln("Fail on get event from PG", err)
			}
			events = append(events, e)
		}
		fmt.Println(events)
	*/

	//IntervalIsBusy
	ev, err := event.CreateEvent("2020-02-29T16:10:00Z", "2020-02-29T18:01:00Z", "Event 1", "Some Desc1")
	if err != nil {
		log.Errorln("Cant parse", err)
	}
	ev.Id = 1

	exist := false
	var err3 error
	var rows *sqlx.Rows
	//SELECT * FROM events WHERE start_time < 'new_end_time' and end_time > 'new_start_time'
	if exist == true {
		rows, err3 = db.NamedQueryContext(ctxExec, "SELECT true FROM events WHERE start_time < :end_time AND end_time > :start_time and id != :id;", ev)
	} else {
		rows, err3 = db.NamedQueryContext(ctxExec, "SELECT true FROM events WHERE start_time < :end_time AND end_time > :start_time and id != :id;", ev)
	}
	if err3 != nil {
		log.Errorln("Fail on check is busy interval", err)
	}
	isBusy := rows.Next()
	fmt.Println(isBusy)

	//Edit(event.Event) error

	/*
		opts := []grpc.DialOption{grpc.WithInsecure()}

		conn, err := grpc.Dial(conf.GrpcListen, opts...)
		if err != nil {
			_, _ = fmt.Fprint(os.Stderr, err)
			os.Exit(2)
		}
		client := protobuf.NewCalendarClient(conn)
		ctx := context.TODO()

		//Event0
		_, err = client.AddEvent(ctx, &protobuf.Event{StartTime: 1000, EndTime: 2000, Title: "title1", Description: "description1"})
		fmt.Println(err)

		//Event1
		_, err = client.AddEvent(ctx, &protobuf.Event{StartTime: 2000, EndTime: 3000, Title: "title2", Description: "description2"})
		fmt.Println(err)

		//Event2
		_, err = client.AddEvent(ctx, &protobuf.Event{StartTime: 1000, EndTime: 4000, Title: "title3", Description: "description3"})
		fmt.Println(err)

		grpcEvent, err := client.GetEvent(ctx, &protobuf.Id{Id: 0})
		fmt.Println(grpcEvent, err)

		grpcEvents, err := client.GetAllEvents(ctx, &empty.Empty{})
		fmt.Println(grpcEvents, err)

		_, _ = client.DelEvent(ctx, &protobuf.Id{Id: 0})

		grpcEvents2, err := client.GetAllEvents(ctx, &empty.Empty{})
		fmt.Println(grpcEvents2, err)
	*/
}
