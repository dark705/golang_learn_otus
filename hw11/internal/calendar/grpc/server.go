package main

import (
	"context"
	"net"

	"github.com/dark705/otus/hw11/internal/protobuf"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

func main() {
	listener, err := net.Listen("tcp", ":5300")

	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}

	opts := []grpc.ServerOption{}
	grpcServer := grpc.NewServer(opts...)

	protobuf.RegisterCalendarServer(grpcServer, &CalendarServer{})

	grpcServer.Serve(listener)
}

type CalendarServer struct{}

func (s *CalendarServer) GetEvent(ctx context.Context, req *protobuf.Id) (*protobuf.Event, error) {
	//TODO!!!!
	return nil, nil
}
