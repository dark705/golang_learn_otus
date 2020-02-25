package grpc

import (
	"context"
	"net"

	"github.com/dark705/otus/hw11/internal/config"
	"github.com/dark705/otus/hw11/internal/protobuf"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/status"
)

func RunServer(conf *config.Config, log *logrus.Logger) {
	log.Info("Start GRPC server:", conf.GrpcListen)

	listener, err := net.Listen("tcp", conf.GrpcListen)
	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	protobuf.RegisterCalendarServer(grpcServer, &CalendarServerGrpc{
		log: log,
	})

	err = grpcServer.Serve(listener)
	if err != nil {
		grpclog.Fatalf("failed to run grpc server: %v", err)
	}

}

type CalendarServerGrpc struct {
	log *logrus.Logger
}

func (s *CalendarServerGrpc) GetEvent(ctx context.Context, req *protobuf.Id) (*protobuf.Event, error) {
	//TODO!!!!
	s.log.Info("Income GerEvent")
	return nil, status.Errorf(codes.Unimplemented, "method GetEvent not implemented")
}
func (s *CalendarServerGrpc) AddEvent(ctx context.Context, req *protobuf.Event) (*protobuf.Event, error) {
	//TODO!!!!
	s.log.Info("Income AddEvent")
	return nil, status.Errorf(codes.Unimplemented, "method AddEvent not implemented")
}
