.PHONY: build
build:
	go build -o ./bin/calendar_scheduler ./cmd/scheduler/scheduler.go
	go build -o ./bin/calendar_sender ./cmd/sender/sender.go
	go build -o ./bin/calendar_api ./cmd/api/api.go
	go build -o ./bin/calendar_api_client ./cmd/api_client/api.client.go
	chmod +x ./bin/calendar_api_client ./bin/calendar_api ./bin/calendar_sender ./bin/calendar_scheduler

test_unit: docker_up_environment  sleep15 go_test docker_down_environment

test: docker_up_environment sleep15 docker_up_app sleep5 docker_run_integration_tests

build_grpc:
	./build/protobuf.grpc.sh

go_test:
	sleep 10 &&\
	go  test -v --count=1 github.com/dark705/otus/hw15/internal/calendar/calendar && \
    go  test -v --count=1 github.com/dark705/otus/hw15/internal/storage && \
    go  test -v --count=1 github.com/dark705/otus/hw15/internal/grpc

docker_up_environment:
	docker-compose -f ./build/docker/docker-compose-environment.yml up -d

docker_down_environment:
	docker-compose -f ./build/docker/docker-compose-environment.yml down

docker_up_app:
	docker-compose -f ./build/docker/docker-compose.yml up -d

docker_down_app:
	docker-compose -f ./build/docker/docker-compose.yml down

docker_run_integration_tests:
	docker-compose -f ./build/docker/docker-compose-tests.yml up

docker_stop_integration_tests:
	docker-compose -f ./build/docker/docker-compose-tests.yml down

sleep15:
	sleep 15

sleep5:
	sleep 5
