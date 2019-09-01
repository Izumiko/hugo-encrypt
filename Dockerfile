FROM golang:alpine AS builder

RUN apk add --no-cache git

RUN git clone https://github.com/Izumiko/hugo-encrypt /hugo-encrypt

WORKDIR /hugo-encrypt

RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/hugo-encrypt

FROM alpine

WORKDIR /data

COPY --from=builder /go/bin/hugo-encrypt /bin

