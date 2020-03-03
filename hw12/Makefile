    run_server:
	go run ./cmd/server.go

    run_client:
	go run ./cmd/client.go

    build_grpc:
	./build/protobuf.grpc.sh

    test: docker_up go_test docker_down

    go_test:
	go  test -v github.com/dark705/otus/hw12/internal/calendar/calendar \
    github.com/dark705/otus/hw12/internal/storage

    docker_up:
	./build/docker.up.sh && sleep 5

    docker_down:
	./build/docker.down.sh

