FROM golang:1.11

RUN mkdir -p /go/src/website
WORKDIR /go/src/website
COPY . /go/src/website

#ONBUILD RUN go get -d -v ./...
ONBUILD RUN go build .

CMD ["./website"]
