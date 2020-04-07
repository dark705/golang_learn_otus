.PHONY: build
build:
	go build -o ./bin/calendar_scheduler ./cmd/scheduler/scheduler.go
	go build -o ./bin/calendar_sender ./cmd/sender/sender.go
	go build -o ./bin/calendar_api ./cmd/api/api.go
	go build -o ./bin/calendar_api_client ./cmd/api_client/api.client.go
	chmod +x ./bin/calendar_api_client ./bin/calendar_api ./bin/calendar_sender ./bin/calendar_scheduler

test_unit: docker_up_environment  sleep15 go_test docker_down_environment

test:
	docker-compose -f ./build/docker/docker-compose-tests.yml up --abort-on-container-exit --exit-code-from godog && \
	docker-compose -f ./build/docker/docker-compose-tests.yml down

build_grpc:
	./build/protobuf.grpc.sh

go_test:
	sleep 10 &&\
	go  test -v --count=1 github.com/dark705/otus/hw16/internal/calendar/calendar && \
    go  test -v --count=1 github.com/dark705/otus/hw16/internal/storage && \
    go  test -v --count=1 github.com/dark705/otus/hw16/internal/grpc

docker_up_environment:
	docker-compose -f ./build/docker/docker-compose-environment.yml up -d

docker_down_environment:
	docker-compose -f ./build/docker/docker-compose-environment.yml down

docker_up_app:
	docker-compose -f ./build/docker/docker-compose.yml up -d

docker_down_app:
	docker-compose -f ./build/docker/docker-compose.yml down

sleep15:
	sleep 15
