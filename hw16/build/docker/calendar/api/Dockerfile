FROM golang:1.13 AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./bin/api ./cmd/api/api.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /app
RUN mkdir config bin logs
COPY --from=builder /app/bin/api ./bin/api
COPY  config/config.docker.yaml ./config/config.yaml
EXPOSE 8081/tcp 5300/tcp 9102/tcp 9103/tcp
CMD ["./bin/api"]
