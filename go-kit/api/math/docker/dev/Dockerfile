FROM golang:1.22

ENV GOPATH /go
ENV GO111MODULE=on
ENV PATH "$PATH:$(go env GOPATH)/bin"

WORKDIR /go/src/app
COPY go.mod go.sum ./

RUN go mod download
RUN go install github.com/cosmtrek/air@latest
RUN go install github.com/go-delve/delve/cmd/dlv@latest
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
