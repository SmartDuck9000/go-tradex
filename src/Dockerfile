FROM golang:1.15-alpine AS build

WORKDIR $GOPATH/src/github.com/SmartDuck9000/go-tradex/src
COPY . .
RUN go get -d -v ./...
RUN go install -v ./...

ENTRYPOINT $GOPATH/bin/stat_service
EXPOSE 5000