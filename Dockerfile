FROM golang:latest
RUN mkdir -p /go/src/app
WORKDIR /go/src/app
RUN apt-get -y update && apt-get -y install git
COPY app/go.mod .
COPY app/go.sum .
RUN go mod download
COPY . .
ENTRYPOINT go run main.go