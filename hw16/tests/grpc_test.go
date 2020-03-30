package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/cucumber/godog"
	"github.com/cucumber/godog/gherkin"
	"github.com/dark705/otus/hw16/pkg/calendar/protobuf"
	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

var grpcError error
var client protobuf.CalendarClient
var sendGrpcEvent protobuf.Event

func iSendExecuteAddEventOnWithData(addr string, data *gherkin.DocString) error {
	ctxConn, _ := context.WithTimeout(context.Background(), time.Second*10)
	conn, err := grpc.DialContext(ctxConn, addr, []grpc.DialOption{grpc.WithInsecure(), grpc.WithBlock()}...)
	if err != nil {
		err = fmt.Errorf("can't connect to grpc server, error: %v", err)
		return err
	}
	client = protobuf.NewCalendarClient(conn)

	replacer := strings.NewReplacer("\n", "", "\t", "")
	cleanJson := replacer.Replace(data.Content)

	err = json.Unmarshal([]byte(cleanJson), &sendGrpcEvent)
	if err != nil {
		err = fmt.Errorf("can't unmarshal, error: %v", err)
		return err
	}

	_, grpcError = client.AddEvent(context.TODO(), &sendGrpcEvent)
	return nil
}

func theResponseErrorCodeShouldBeNil() (err error) {
	if grpcError != nil {
		err = fmt.Errorf("fail on execute AddEvent %v", grpcError)
	}
	return err
}

func theExecuteGetEventReturnSameEventWithNoErrorCode() error {
	getGrpcEvent, err := client.GetEvent(context.TODO(), &protobuf.Id{Id: 1})
	if err != nil {
		return err
	}

	if !proto.Equal(&sendGrpcEvent, getGrpcEvent) {
		return errors.New("add and Get Event's not same")
	}
	return nil
}

func afterExecuteDelEventReturnNoErrorCode() error {
	_, err := client.DelEvent(context.TODO(), &protobuf.Id{Id: 1})
	if err != nil {
		return err
	}
	return nil
}

func iSendExecuteAddEventOnTwoTimesWithData(addr string, data *gherkin.DocString) error {
	ctxConn, _ := context.WithTimeout(context.Background(), time.Second*10)
	conn, err := grpc.DialContext(ctxConn, addr, []grpc.DialOption{grpc.WithInsecure(), grpc.WithBlock()}...)
	if err != nil {
		err = fmt.Errorf("can't connect to grpc server, error: %v", err)
		return err
	}
	client = protobuf.NewCalendarClient(conn)

	replacer := strings.NewReplacer("\n", "", "\t", "")
	cleanJson := replacer.Replace(data.Content)

	err = json.Unmarshal([]byte(cleanJson), &sendGrpcEvent)
	if err != nil {
		err = fmt.Errorf("can't unmarshal, error: %v", err)
		return err
	}

	_, err = client.AddEvent(context.TODO(), &sendGrpcEvent)
	if err != nil {
		return err
	}
	_, grpcError = client.AddEvent(context.TODO(), &sendGrpcEvent)
	return nil
}

func theResponseErrorDescShouldBe(desc string) error {
	s, ok := status.FromError(grpcError)
	if !ok {
		return errors.New("can't get status from grpcError")
	}
	if desc != s.Message() {
		return errors.New("wrong desc on error busy")
	}
	return nil
}

func FeatureGrpcContext(s *godog.Suite) {
	s.Step(`^I send execute AddEvent on "([^"]*)" with data:$`, iSendExecuteAddEventOnWithData)
	s.Step(`^The response error code should be nil$`, theResponseErrorCodeShouldBeNil)
	s.Step(`^The execute GetEvent return same Event, with no error code$`, theExecuteGetEventReturnSameEventWithNoErrorCode)
	s.Step(`^After Execute DelEvent, return no error code$`, afterExecuteDelEventReturnNoErrorCode)
	s.Step(`^I send execute AddEvent on "([^"]*)" two times with data:$`, iSendExecuteAddEventOnTwoTimesWithData)
	s.Step(`^The response error desc should be "([^"]*)"$`, theResponseErrorDescShouldBe)
}
