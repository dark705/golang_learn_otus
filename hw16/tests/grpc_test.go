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
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

var ctx context.Context
var grpcError error
var client protobuf.CalendarClient
var sendGrpcEvent protobuf.Event

func iSendExecuteAddEventOnEventData(addr string, data *gherkin.DocString) error {
	ctxConn, _ := context.WithTimeout(context.Background(), time.Second*2)
	conn, err := grpc.DialContext(ctxConn, addr, []grpc.DialOption{grpc.WithInsecure(), grpc.WithBlock()}...)
	if err != nil {
		err = fmt.Errorf("can't connect to grpc server, error: %v", err)
		return err
	}

	replacer := strings.NewReplacer("\n", "", "\t", "")
	cleanJson := replacer.Replace(data.Content)

	err = json.Unmarshal([]byte(cleanJson), &sendGrpcEvent)
	if err != nil {
		err = fmt.Errorf("can't unmarshal, error: %v", err)
		return err
	}
	fmt.Println(cleanJson)

	client = protobuf.NewCalendarClient(conn)
	ctx = context.TODO()
	_, grpcError = client.AddEvent(ctx, &sendGrpcEvent)
	return grpcError
}

func theResponseErrorCodeShouldBeNil() (err error) {
	if grpcError != nil {
		err = fmt.Errorf("fail on execute AddEvent %v", grpcError)
	}
	return err
}

func theExecuteGetAllEventsReturnOneSameEvent() error {
	grpcEvents, err := client.GetAllEvents(ctx, &empty.Empty{})
	if err != nil {
		return err
	}
	if len(grpcEvents.Events) != 1 {
		return errors.New("grpc client returned not one Event")
	}

	if !proto.Equal(&sendGrpcEvent, grpcEvents.Events[0]) {
		fmt.Println(&sendGrpcEvent)
		fmt.Println(grpcEvents.Events[0])
		return errors.New("add and Get Event's not same")
	}
	return nil
}

func FeatureGrpcContext(s *godog.Suite) {
	s.Step(`^I send execute AddEvent on "([^"]*)" Event data:$`, iSendExecuteAddEventOnEventData)
	s.Step(`^The response error code should be nil$`, theResponseErrorCodeShouldBeNil)
	s.Step(`^The execute GetAllEvents return one same Event$`, theExecuteGetAllEventsReturnOneSameEvent)

}
