package grpc

import (
	"context"
	"net"
	"time"

	"github.com/dark705/otus/hw14/internal/calendar/calendar"
	"github.com/dark705/otus/hw14/internal/calendar/event"
	"github.com/dark705/otus/hw14/pkg/calendar/protobuf"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Config struct {
	GrpcListen string
}

type Server struct {
	Config   Config
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
	if err != nil {
		s.Logger.Errorln("Error on Calendar, add:", err)
		if err == calendar.ErrDateBusy {
			return &empty.Empty{}, status.Error(codes.AlreadyExists, "Date interval for new event already busy")
		}
		return &empty.Empty{}, status.Error(codes.Internal, "Internal server error")
	}

	return &empty.Empty{}, nil
}

func (s *Server) GetEvent(ctx context.Context, grpcId *protobuf.Id) (*protobuf.Event, error) {
	s.Logger.Debug("Income gRPC GetEvent() id:", grpcId)
	calendarEvent, err := s.Calendar.GetEvent(int(grpcId.Id))
	if err != nil {
		s.Logger.Errorln("Error on Calendar, get:", err)
		return &protobuf.Event{}, status.Error(codes.NotFound, "Even not found")
	}

	return &protobuf.Event{
		Id:          int32(calendarEvent.Id),
		StartTime:   calendarEvent.StartTime.Unix(),
		EndTime:     calendarEvent.EndTime.Unix(),
		Title:       calendarEvent.Title,
		Description: calendarEvent.Description,
	}, nil
}

func (s *Server) DelEvent(ctx context.Context, grpcId *protobuf.Id) (*empty.Empty, error) {
	s.Logger.Debug("Income gRPC DelEvent() id:", grpcId)
	err := s.Calendar.DelEvent(int(grpcId.Id))
	if err != nil {
		s.Logger.Errorln("Error on Calendar, delete:", err)
		return &empty.Empty{}, status.Error(codes.NotFound, "Even not found")
	}

	return &empty.Empty{}, nil
}

func (s *Server) EditEvent(ctx context.Context, grpcE *protobuf.Event) (*empty.Empty, error) {
	s.Logger.Debug("Income gRPC EditEvent() event:", grpcE)
	err := s.Calendar.EditEvent(event.Event{
		StartTime:   time.Unix(grpcE.StartTime, 0),
		EndTime:     time.Unix(grpcE.EndTime, 0),
		Title:       grpcE.Title,
		Description: grpcE.Description,
	})
	if err != nil {
		s.Logger.Errorln("Error on Calendar, edit:", err)
		if err == calendar.ErrDateBusy {
			return &empty.Empty{}, status.Error(codes.AlreadyExists, "Date interval for new event already busy")
		}
		return &empty.Empty{}, status.Error(codes.Internal, "Internal server error")
	}

	return &empty.Empty{}, nil
}

func (s *Server) GetAllEvents(ctx context.Context, ev *empty.Empty) (*protobuf.Events, error) {
	s.Logger.Debug("Income gRPC GetAllEvents()")
	calendarEvents, err := s.Calendar.GetAllEvents()
	l := len(calendarEvents)
	protobufEvents := make([]*protobuf.Event, 0, l)
	if err != nil {
		s.Logger.Errorln("Error to get all Events calendar", err)
		return &protobuf.Events{Events: protobufEvents}, status.Error(codes.Internal, "Internal server error")
	}

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

	return &protobuf.Events{Events: protobufEvents}, nil
}
