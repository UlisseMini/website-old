FROM golang:1.11

RUN mkdir -p /go/src/website/resources
WORKDIR /go/src/website
COPY . /go/src/website

RUN go build .

CMD ["./website"]
