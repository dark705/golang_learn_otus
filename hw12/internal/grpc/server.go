package grpc

import (
	"context"
	"net"
	"time"

	"github.com/dark705/otus/hw11/internal/calendar/calendar"
	"github.com/dark705/otus/hw11/internal/calendar/event"
	"github.com/dark705/otus/hw11/internal/config"
	"github.com/dark705/otus/hw11/pkg/calendar/protobuf"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type Server struct {
	Config   config.Config
	Logger   *logrus.Logger
	Calendar *calendar.Calendar
	server   *grpc.Server
}

func (s *Server) Run() {
	s.Logger.Info("Start GRPC server:", s.Config.GrpcListen)

	listener, err := net.Listen("tcp", s.Config.GrpcListen)
	if err != nil {
		s.Logger.Fatalf("failed to listen: %v", err)
	}

	s.server = grpc.NewServer()
	protobuf.RegisterCalendarServer(s.server, s)

	err = s.server.Serve(listener)
	if err != nil {
		s.Logger.Fatalf("failed to run grpc server: %v", err)
	}
}

func (s *Server) Shutdown() {
	s.Logger.Info("Graceful shutdown GRPC server...")
	s.server.GracefulStop()
}

type CalendarServerGrpc struct {
	log      *logrus.Logger
	calendar *calendar.Calendar
}

func (s *Server) AddEvent(ctx context.Context, grpcE *protobuf.Event) (*empty.Empty, error) {
	s.Logger.Debug("Income gRPC AddEvent() event: ", grpcE)
	err := s.Calendar.AddEvent(event.Event{
		StartTime:   time.Unix(grpcE.StartTime, 0),
		EndTime:     time.Unix(grpcE.EndTime, 0),
		Title:       grpcE.Title,
		Description: grpcE.Description,
	})

	return &empty.Empty{}, err
}

func (s *Server) GetEvent(ctx context.Context, grpcId *protobuf.Id) (*protobuf.Event, error) {
	s.Logger.Debug("Income gRPC GetEvent() id:", grpcId)
	calendarEvent, err := s.Calendar.GetEvent(int(grpcId.Id))

	return &protobuf.Event{
		Id:          int32(calendarEvent.Id),
		StartTime:   calendarEvent.StartTime.Unix(),
		EndTime:     calendarEvent.EndTime.Unix(),
		Title:       calendarEvent.Title,
		Description: calendarEvent.Description,
	}, err
}

func (s *Server) DelEvent(ctx context.Context, grpcId *protobuf.Id) (*empty.Empty, error) {
	s.Logger.Debug("Income gRPC DelEvent() id:", grpcId)

	return &empty.Empty{}, s.Calendar.DelEvent(int(grpcId.Id))
}

func (s *Server) EditEvent(ctx context.Context, grpcE *protobuf.Event) (*empty.Empty, error) {
	s.Logger.Debug("Income gRPC EditEvent() event:", grpcE)

	err := s.Calendar.EditEvent(event.Event{
		StartTime:   time.Unix(grpcE.StartTime, 0),
		EndTime:     time.Unix(grpcE.EndTime, 0),
		Title:       grpcE.Title,
		Description: grpcE.Description,
	})

	return &empty.Empty{}, err
}

func (s *Server) GetAllEvents(ctx context.Context, ev *empty.Empty) (*protobuf.Events, error) {
	s.Logger.Debug("Income gRPC GetAllEvents()")
	calendarEvents, err := s.Calendar.GetAllEvents()
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
