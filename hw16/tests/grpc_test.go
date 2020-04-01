package tests

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/cucumber/godog"
	"github.com/cucumber/godog/gherkin"
	"github.com/dark705/otus/hw16/internal/helpers"
	"github.com/dark705/otus/hw16/pkg/calendar/protobuf"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

var grpcError error
var client protobuf.CalendarClient
var sendGrpcEvent protobuf.Event
var lastEventID int32

func iSendExecuteAddEventWithData(data *gherkin.DocString) (err error) {

	replacer := strings.NewReplacer("\n", "", "\t", "")
	cleanJSON := replacer.Replace(data.Content)

	err = json.Unmarshal([]byte(cleanJSON), &sendGrpcEvent)
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
	getGrpcEvents, err := client.GetAllEvents(context.TODO(), &empty.Empty{})
	lastEventID = getGrpcEvents.Events[len(getGrpcEvents.Events)-1].Id

	getGrpcEvent, err := client.GetEvent(context.TODO(), &protobuf.Id{Id: lastEventID})
	if err != nil {
		return err
	}
	sendGrpcEvent.Id = lastEventID

	if !proto.Equal(&sendGrpcEvent, getGrpcEvent) {
		return errors.New("add and Get Event's not same")
	}
	return nil
}

func afterExecuteDelEventReturnNoErrorCode() error {
	_, err := client.DelEvent(context.TODO(), &protobuf.Id{Id: lastEventID})
	if err != nil {
		return err
	}
	return nil
}

func iSendExecuteAddEventTwoTimesWithData(data *gherkin.DocString) (err error) {

	replacer := strings.NewReplacer("\n", "", "\t", "")
	cleanJSON := replacer.Replace(data.Content)

	err = json.Unmarshal([]byte(cleanJSON), &sendGrpcEvent)
	if err != nil {
		err = fmt.Errorf("can't unmarshal, error: %v", err)
		return err
	}

	_, err = client.AddEvent(context.TODO(), &sendGrpcEvent)
	if err != nil {
		return err
	}
	_, grpcError = client.AddEvent(context.TODO(), &sendGrpcEvent)

	getGrpcEvents, err := client.GetAllEvents(context.TODO(), &empty.Empty{})
	lastEventID = getGrpcEvents.Events[len(getGrpcEvents.Events)-1].Id

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

func initGRPCClient() {
	ctxConn, _ := context.WithTimeout(context.Background(), time.Second*10)
	conn, err := grpc.DialContext(ctxConn, "calendar_api:5300", []grpc.DialOption{grpc.WithInsecure(), grpc.WithBlock()}...)
	helpers.FailOnError(err, "gRPC client fail")
	client = protobuf.NewCalendarClient(conn)
}

func FeatureGrpcContext(s *godog.Suite) {
	initGRPCClient()

	s.Step(`^I send execute AddEvent with data:$`, iSendExecuteAddEventWithData)
	s.Step(`^The response error code should be nil$`, theResponseErrorCodeShouldBeNil)
	s.Step(`^The execute GetEvent return same Event, with no error code$`, theExecuteGetEventReturnSameEventWithNoErrorCode)
	s.Step(`^After Execute DelEvent, return no error code$`, afterExecuteDelEventReturnNoErrorCode)
	s.Step(`^I send execute AddEvent two times with data:$`, iSendExecuteAddEventTwoTimesWithData)
	s.Step(`^The response error desc should be "([^"]*)"$`, theResponseErrorDescShouldBe)
	s.Step(`^After Execute DelEvent, return no error code$`, afterExecuteDelEventReturnNoErrorCode)
}
