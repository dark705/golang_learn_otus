package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	"github.com/dark705/otus/hw15/internal/config"
	"github.com/dark705/otus/hw15/internal/logger"
	"github.com/dark705/otus/hw15/pkg/calendar/protobuf"
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

	log := logger.NewLogger(logger.Config{
		File:  conf.Logger.File,
		Level: conf.Logger.Level,
	})
	defer logger.CloseLogFile()

	ctxConn, _ := context.WithTimeout(context.Background(), time.Second*2)
	conn, err := grpc.DialContext(ctxConn, conf.Api.GrpcListen, []grpc.DialOption{grpc.WithInsecure(), grpc.WithBlock()}...)
	if err != nil {
		log.Fatalln("Can't connect to grpc server, error: ", err)
	}

	client := protobuf.NewCalendarClient(conn)
	ctx := context.TODO()

	//Event
	_, err = client.AddEvent(ctx, &protobuf.Event{StartTime: 1000, EndTime: 2000, Title: "title1", Description: "description1"})
	LogOnError(log, "Fail on add Event", err)

	//Event
	_, err = client.AddEvent(ctx, &protobuf.Event{StartTime: 2000, EndTime: 3000, Title: "title2", Description: "description2"})
	LogOnError(log, "Fail on add Event", err)

	//Event
	_, err = client.AddEvent(ctx, &protobuf.Event{StartTime: 3000, EndTime: 4000, Title: "title3", Description: "description3"})
	LogOnError(log, "Fail on add Event", err)

	//getAllEvents
	grpcEvents, err := client.GetAllEvents(ctx, &empty.Empty{})
	LogOnError(log, "Fail on get all Event's", err)
	if err == nil {
		fmt.Println(grpcEvents)
		lastId := grpcEvents.Events[len(grpcEvents.Events)-1].Id

		_, err = client.DelEvent(ctx, &protobuf.Id{Id: lastId})
		LogOnError(log, "Fail on del Event's", err)

		_, err = client.DelEvent(ctx, &protobuf.Id{Id: lastId - 1})

		LogOnError(log, "Fail on del Event's", err)

		_, err = client.DelEvent(ctx, &protobuf.Id{Id: lastId - 2})
		LogOnError(log, "Fail on del Event's", err)
	}

}

func LogOnError(log logrus.Logger, mes string, err error) {
	if err != nil {
		log.Errorln(fmt.Sprintf("%s, error: %v", mes, err))
	}
}
