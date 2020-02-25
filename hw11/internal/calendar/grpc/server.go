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
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/status"
)

func RunServer(conf *config.Config, log *logrus.Logger, calendar *calendar.Calendar) {
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

func (s *CalendarServerGrpc) GetEvent(ctx context.Context, id *protobuf.Id) (*protobuf.Event, error) {
	//TODO!!!!
	s.log.Info("Income GerEvent")
	return nil, status.Errorf(codes.Unimplemented, "method GetEvent not implemented")
}
func (s *CalendarServerGrpc) AddEvent(ctx context.Context, ev *protobuf.Event) (*empty.Empty, error) {
	s.log.Debug("Income gRPC: AddEvent:", ev)
	err := s.calendar.AddEvent(event.Event{
		StartTime:   time.Unix(ev.StartTime, 0),
		EndTime:     time.Unix(ev.EndTime, 0),
		Title:       ev.Title,
		Description: ev.Description,
	})
	return &empty.Empty{}, err
}
