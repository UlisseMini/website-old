# First stage, build the program
FROM golang as builder
WORKDIR /build

# Ensure go mod dependences
ENV GO111MODULE=on

COPY go.mod .
# uncomment this if you've got dependences.
#COPY go.sum .

RUN go mod download

# Build the program
COPY . .
RUN CGO_ENABLED=0 go build -o binary

# Second stage, take only the binary
FROM scratch
COPY --from=builder /build/binary /binary

ENTRYPOINT ["./binary"]
