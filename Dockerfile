FROM golang:1.11

ONBUILD RUN mkdir -p /go/src/website
ONBUILD RUN mkdir -p /etc/letsencrypt/archive/gopher.ddns.net/
ONBUILD WORKDIR /go/src/website
ONBUILD COPY . /go/src/website
ONBUILD COPY /etc/letsencrypt/archive/gopher.ddns.net/cert1.pem /etc/letsencrypt/archive/gopher.ddns.net
ONBUILD COPY /etc/letsencrypt/archive/gopher.ddns.net/privkey1.pem /etc/letsencrypt/archive/gopher.ddns.net

ONBUILD RUN go build .

CMD ["./website"]
