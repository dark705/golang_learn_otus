build: build_scheduler build_sender build_api build_api_client

test: docker_up go_test docker_down

build_scheduler:
	go build -o ./bin/calendar_scheduler ./cmd/scheduler.go
	chmod +x ./cmd/scheduler.go

build_sender:
	go build -o ./bin/calendar_sender ./cmd/sender.go
	chmod +x ./cmd/sender.go

build_api:
	go build -o ./bin/calendar_api ./cmd/api.go
	chmod +x ./cmd/api.go

build_api_client:
	go build -o ./bin/calendar_api_client ./cmd/api.client.go
	chmod +x ./cmd/api.client.go

run_scheduler:
	go run ./cmd/scheduler.go

run_sender:
	go run ./cmd/sender.go

run_api:
	go run ./cmd/api.go

run_client_api:
	go run ./cmd/api.client.go

build_grpc:
	./build/protobuf.grpc.sh

go_test:
	go  test -v --count=1 github.com/dark705/otus/hw14/internal/calendar/calendar && \
    go  test -v --count=1 github.com/dark705/otus/hw14/internal/storage && \
    go  test -v --count=1 github.com/dark705/otus/hw14/internal/grpc

docker_up:
	./build/docker.up.sh && sleep 5

docker_down:
	./build/docker.down.sh

