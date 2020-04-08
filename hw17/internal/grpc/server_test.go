package grpc

import (
	"context"
	"testing"
	"time"

	"github.com/dark705/otus/hw17/internal/calendar/calendar"
	"github.com/dark705/otus/hw17/internal/storage"
	"github.com/dark705/otus/hw17/pkg/calendar/protobuf"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func TestAddEventGetEvent(t *testing.T) {
	inMemory := storage.InMemory{}
	err := inMemory.Init()
	if err != nil {
		t.Error("Can't init storage")
	}
	cal := calendar.Calendar{Storage: &inMemory, Logger: &logrus.Logger{}}
	grpcServer := Server{Config: Config{GrpcListen: "127.0.0.1:53001", PrometheusListen: "127.0.0.1:53002"}, Calendar: &cal, Logger: &logrus.Logger{}}
	ctx := context.Background()
	go grpcServer.Run()
	time.Sleep(time.Second) // wait for grpc server run
	defer grpcServer.Shutdown()

	conn, err := grpc.Dial("127.0.0.1:53001", []grpc.DialOption{grpc.WithInsecure()}...)
	if err != nil {
		t.Error("Fail connect to GRPC server")
	}

	client := protobuf.NewCalendarClient(conn)

	sendGrpcEvent := protobuf.Event{Id: 1, StartTime: 1000, EndTime: 2000, Title: "title1", Description: "description1"}
	_, err = client.AddEvent(ctx, &sendGrpcEvent)
	if err != nil {
		t.Error("Fail AddEvent")
	}

	getGrpcEvent, err := client.GetEvent(ctx, &protobuf.Id{Id: 1})
	if err != nil {
		t.Error("Fail GetEvent() with id=1")
	}
	if !proto.Equal(&sendGrpcEvent, getGrpcEvent) {
		t.Error("Add and Get Event's not same")
	}
}

func TestDelGetAllEvents(t *testing.T) {
	inMemory := storage.InMemory{}
	err := inMemory.Init()
	if err != nil {
		t.Error("Can't init storage")
	}
	cal := calendar.Calendar{Storage: &inMemory, Logger: &logrus.Logger{}}
	grpcServer := Server{Config: Config{GrpcListen: "127.0.0.1:53001", PrometheusListen: "127.0.0.1:53002"}, Calendar: &cal, Logger: &logrus.Logger{}}
	ctx := context.Background()
	go grpcServer.Run()
	time.Sleep(time.Second) // wait for grpc server run
	defer grpcServer.Shutdown()

	conn, err := grpc.Dial("127.0.0.1:53001", []grpc.DialOption{grpc.WithInsecure()}...)
	if err != nil {
		t.Error("Fail connect to GRPC server")
	}

	client := protobuf.NewCalendarClient(conn)

	_, err = client.AddEvent(ctx, &protobuf.Event{StartTime: 1000, EndTime: 2000, Title: "title1", Description: "description1"})
	_, err = client.AddEvent(ctx, &protobuf.Event{StartTime: 2000, EndTime: 3000, Title: "title2", Description: "description2"})
	_, err = client.AddEvent(ctx, &protobuf.Event{StartTime: 4000, EndTime: 5000, Title: "title3", Description: "description3"})
	_, err = client.AddEvent(ctx, &protobuf.Event{StartTime: 6000, EndTime: 7000, Title: "title4", Description: "description4"})
	if err != nil {
		t.Error("Fail AddEvent 4 Events")
	}

	_, err = client.DelEvent(ctx, &protobuf.Id{Id: 1})
	if err != nil {
		t.Error("Fail DelEvent() with id=1")
	}

	_, err = client.DelEvent(ctx, &protobuf.Id{Id: 2})
	if err != nil {
		t.Error("Fail DelEvent() with id=2")
	}

	getGrpcEvents, err := client.GetAllEvents(ctx, &empty.Empty{})
	if err != nil {
		t.Error("Fail get GetAllEvents()")
	}

	if len(getGrpcEvents.Events) != 2 {
		t.Error("Len of GetAllEvents() not return 2, after AddEvent() 4 times and DelEvent() 2 times")
	}
}
