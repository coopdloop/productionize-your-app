FROM golang:1.22-alpine3.19 AS builder

# installing git
RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh

# setting working directory
WORKDIR /go/src/app

# installing dependencies
COPY go.mod ./
RUN go mod download

COPY / /go/src/app/
RUN go build -o myapp

FROM alpine:latest

RUN apk update && apk upgrade && \
    apk add --no-cache openssh curl ca-certificates

WORKDIR /go/src/app
COPY --from=builder /go/src/app/myapp /go/src/app/myapp

EXPOSE 5000

CMD ["./myapp"]
