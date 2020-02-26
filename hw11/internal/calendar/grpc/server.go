package grpc

import (
	"context"
	"net"
	"time"

	"github.com/dark705/otus/hw11/internal/calendar/calendar"
	"github.com/dark705/otus/hw11/internal/calendar/event"
	"github.com/dark705/otus/hw11/internal/config"
	"github.com/dark705/otus/hw11/internal/protobuf"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

func RunServer(conf config.Config, log *logrus.Logger, calendar *calendar.Calendar) {
	log.Info("Start GRPC server:", conf.GrpcListen)

	listener, err := net.Listen("tcp", conf.GrpcListen)
	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	protobuf.RegisterCalendarServer(grpcServer, &CalendarServerGrpc{
		log:      log,
		calendar: calendar,
	})

	err = grpcServer.Serve(listener)
	if err != nil {
		grpclog.Fatalf("failed to run grpc server: %v", err)
	}

}

type CalendarServerGrpc struct {
	log      *logrus.Logger
	calendar *calendar.Calendar
}

func (s *CalendarServerGrpc) AddEvent(ctx context.Context, grpcE *protobuf.Event) (*empty.Empty, error) {
	s.log.Debug("Income gRPC AddEvent() event: ", grpcE)
	err := s.calendar.AddEvent(event.Event{
		StartTime:   time.Unix(grpcE.StartTime, 0),
		EndTime:     time.Unix(grpcE.EndTime, 0),
		Title:       grpcE.Title,
		Description: grpcE.Description,
	})

	return &empty.Empty{}, err
}

func (s *CalendarServerGrpc) GetEvent(ctx context.Context, grpcId *protobuf.Id) (*protobuf.Event, error) {
	s.log.Debug("Income gRPC GetEvent() id:", grpcId)
	calendarEvent, err := s.calendar.GetEvent(int(grpcId.Id))

	return &protobuf.Event{
		Id:          int32(calendarEvent.Id),
		StartTime:   calendarEvent.StartTime.Unix(),
		EndTime:     calendarEvent.EndTime.Unix(),
		Title:       calendarEvent.Title,
		Description: calendarEvent.Description,
	}, err
}

func (s *CalendarServerGrpc) DelEvent(ctx context.Context, grpcId *protobuf.Id) (*empty.Empty, error) {
	s.log.Debug("Income gRPC DelEvent() id:", grpcId)

	return &empty.Empty{}, s.calendar.DelEvent(int(grpcId.Id))
}

func (s *CalendarServerGrpc) EditEvent(ctx context.Context, grpcE *protobuf.Event) (*empty.Empty, error) {
	s.log.Debug("Income gRPC EditEvent() event:", grpcE)

	s.log.Debug("Income gRPC AddEvent() event: ", grpcE)
	err := s.calendar.EditEvent(event.Event{
		StartTime:   time.Unix(grpcE.StartTime, 0),
		EndTime:     time.Unix(grpcE.EndTime, 0),
		Title:       grpcE.Title,
		Description: grpcE.Description,
	})

	return &empty.Empty{}, err
}

func (s *CalendarServerGrpc) GetAllEvents(ctx context.Context, ev *empty.Empty) (*protobuf.Events, error) {
	s.log.Debug("Income gRPC GetAllEvents()")
	calendarEvents, err := s.calendar.GetAllEvents()
	l := len(calendarEvents)

	protobufEvents := make([]*protobuf.Event, 0, l)

	for _, calendarEvent := range calendarEvents {
		protobufEvent := protobuf.Event{
			Id:          int32(calendarEvent.Id),
			StartTime:   calendarEvent.StartTime.Unix(),
			EndTime:     calendarEvent.EndTime.Unix(),
			Title:       calendarEvent.Title,
			Description: calendarEvent.Description,
		}
		protobufEvents = append(protobufEvents, &protobufEvent)
	}

	return &protobuf.Events{Events: protobufEvents}, err
}
