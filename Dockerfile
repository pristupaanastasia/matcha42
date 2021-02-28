FROM golang:latest
RUN mkdir -p /go/src/app
WORKDIR /go/src/app
RUN apt-get -y update && apt-get -y install git
RUN go get github.com/lib/pq
RUN go get github.com/google/uuid
RUN go get golang.org/x/net/context
ENTRYPOINT go run main.go