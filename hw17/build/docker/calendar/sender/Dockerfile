FROM golang:1.13 AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./bin/sender ./cmd/sender/sender.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /app
RUN mkdir config bin logs
COPY --from=builder /app/bin/sender ./bin/sender
COPY  config/config.docker.yaml ./config/config.yaml
EXPOSE 9104/tcp
CMD ["./bin/sender"]
