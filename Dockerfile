FROM golang:latest
RUN go get -u github.com/oxeque/realize
RUN mkdir /go/src/work
WORKDIR /go/src/work
ADD . /go/src/work