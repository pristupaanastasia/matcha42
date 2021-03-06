FROM golang:latest
RUN mkdir -p /go/src/app
WORKDIR /go/src/app
RUN apt-get -y update && apt-get -y install git
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o main .
ENTRYPOINT go run main.go