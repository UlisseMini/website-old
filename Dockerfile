FROM golang:latest

WORKDIR /go/github.com/itsyourboychipsahoy/website
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["website"]
