package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"time"

	_ "github.com/jackc/pgx/stdlib"

	"github.com/dark705/otus/hw12/internal/calendar/event"
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

	ctxExec, _ := context.WithTimeout(context.Background(), time.Second*time.Duration(conf.PgTimeoutExecute))
	rows, err := db.QueryxContext(ctxExec, "select * from events")
	if err != nil {
		log.Errorln("Fail on db quiery:")
	} else {
		for rows.Next() {
			var e event.Event
			if err := rows.StructScan(&e); err != nil {
				//TODO
			}
			fmt.Println(e)
		}
	}
	fmt.Println("************")
	events := []event.Event{}
	err = db.Select(&events, "select * from events")
	fmt.Println(events, err)

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
