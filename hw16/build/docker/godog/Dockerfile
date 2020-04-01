FROM golang:1.13
WORKDIR /app
RUN GO111MODULE=on go get github.com/cucumber/godog/cmd/godog@v0.8.1
COPY . .
