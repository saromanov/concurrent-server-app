FROM golang:1.11

ADD . /go/app
RUN go build .
ENTRYPOINT main