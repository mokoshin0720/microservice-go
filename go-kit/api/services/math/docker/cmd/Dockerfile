FROM golang:1.22

ENV GOPATH /go
ENV GO111MODULE=on

WORKDIR /go/src/app
COPY . .

ENTRYPOINT [ "go", "run" ]
