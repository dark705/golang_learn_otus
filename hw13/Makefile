    build_scheduler:
	go build -o ./bin/sheduler ./cmd/scheduler.go

    build_sender:
	go build -o ./bin/sender ./cmd/sender.go

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

    test: docker_up go_test docker_down

    go_test:
	go  test -v --count=1 github.com/dark705/otus/hw13/internal/calendar/calendar && \
    go  test -v --count=1 github.com/dark705/otus/hw13/internal/storage && \
    go  test -v --count=1 github.com/dark705/otus/hw13/internal/grpc

    docker_up:
	./build/docker.up.sh && sleep 5

    docker_down:
	./build/docker.down.sh

